package cache

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/syncx"
)

var (
	onceNewRedigo    syncx.Once
	onceNewRedigoRes *Redigo
	onceNewRedigoErr error
)

// NewRedigo return a redis client.
func NewRedigo(config *RedisConfiguration) (*Redigo, error) {
	onceNewRedigo.Do(func() {
		// Default configuration for max active and wait.
		if config.OpenConnectionLimit == 0 &&
			!config.WaitOpenConnection {
			config.OpenConnectionLimit = 5
			config.WaitOpenConnection = true
		}

		if len(config.Addresses) == 0 {
			onceNewRedigoErr = errorx.E("missing address", errorx.CodeConfig)
			return
		}

		// Create connection client.
		client := &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", config.Addresses[0])
			},
			DialContext: nil,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:         int(config.IdleConnectionLimit),
			MaxActive:       int(config.OpenConnectionLimit),
			IdleTimeout:     0,
			Wait:            config.WaitOpenConnection,
			MaxConnLifetime: 0,
		}

		// Try to dial the redis.
		// On error close previous open connection client.
		if _, err := client.Dial(); err != nil {
			_ = client.Close()
			onceNewRedigoErr = errorx.E(err, errorx.CodeGateway)
			return
		}

		onceNewRedigoRes = &Redigo{
			client: client,
		}
	})

	return onceNewRedigoRes, onceNewRedigoErr
}

// Redigo returns a redis client using redigo library.
type Redigo struct {
	client *redis.Pool
}

// Get gets the value from redis in []byte form.
func (r *Redigo) Get(_ context.Context, key string) ([]byte, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "GET"
	data, err := redis.Bytes(con.Do(commandName, key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return data, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// SimpleSet sets value to key in redis without any additional options.
// Key doesn't have a TTL.
func (r *Redigo) SimpleSet(_ context.Context, key, value string) error {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "SET"
	data, err := redis.String(con.Do(commandName, key, value))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return errorx.E(err, errorx.CodeGateway)
	}

	// Extra check for set operations.
	if errors.Is(err, redis.ErrNil) || !strings.EqualFold("OK", data) {
		return errorx.E("redis operation ended unsuccessfully", errorx.CodeGateway)
	}

	return nil
}

// SetEX sets the value to a key with timeout in seconds.
func (r *Redigo) SetEX(_ context.Context, key string, seconds int64, value string) error {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "SET"
	data, err := redis.String(con.Do(commandName, key, value, "ex", seconds))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return errorx.E(err, errorx.CodeGateway)
	}

	// Extra check for set operations.
	if errors.Is(err, redis.ErrNil) || !strings.EqualFold("OK", data) {
		return errorx.E("redis operation ended unsuccessfully", errorx.CodeGateway)
	}

	return nil
}

// SetNX sets a value to a key with specified timeouts.
// SetNX returns false if the key exists.
func (r *Redigo) SetNX(
	_ context.Context,
	key string,
	seconds int64,
	value string,
) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "SET"
	data, err := redis.String(con.Do(commandName, key, value, "ex", seconds, "nx"))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return false, errorx.E(err, errorx.CodeGateway)
	}
	// extra check for set operations
	if errors.Is(err, redis.ErrNil) || !strings.EqualFold("OK", data) {
		return false, nil
	}

	return true, nil
}

// HMGet gets a value of multiple fields from hash key.
func (r *Redigo) HMGet(_ context.Context, key string, fields ...string) ([][]byte, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "HMGET"
	data, err := redis.ByteSlices(con.Do(commandName, key, fields))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return data, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// Exists checks whether the key exists in redis.
func (r *Redigo) Exists(_ context.Context, key string) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "EXISTS"
	data, err := redis.Int64(con.Do(commandName, key))
	if err != nil {
		return false, errorx.E(err, errorx.CodeGateway)
	}
	if data != 1 {
		return false, nil
	}

	return true, nil
}

// Expire sets the ttl of a key to specified value in seconds.
func (r *Redigo) Expire(_ context.Context, key string, seconds int64) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "EXPIRE"
	data, err := redis.Int64(con.Do(commandName, key, seconds))
	if err != nil {
		return false, errorx.E(err, errorx.CodeGateway)
	}

	return data == 1, nil
}

// ExpireAt sets the ttl of a key to a certain timestamp.
func (r *Redigo) ExpireAt(_ context.Context, key string, timestamp int64) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "EXPIREAT"
	data, err := redis.Int64(con.Do(commandName, key, timestamp))
	if err != nil {
		return false, errorx.E(err, errorx.CodeGateway)
	}
	if data != 1 {
		return false, nil
	}

	return true, nil
}

// Incr increments the integer value of a key by 1.
func (r *Redigo) Incr(_ context.Context, key string) (int64, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "INCR"
	data, err := redis.Int64(con.Do(commandName, key))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// Decr decrements the integer value of a key by 1.
func (r *Redigo) Decr(_ context.Context, key string) (int64, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "DECR"
	data, err := redis.Int64(con.Do(commandName, key))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// TTL gets the time to live of a key / expiry time.
func (r *Redigo) TTL(_ context.Context, key string) (int64, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "TTL"
	data, err := redis.Int64(con.Do(commandName, key))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// HGet gets the value of a hash field.
func (r *Redigo) HGet(_ context.Context, key, field string) ([]byte, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "HGET"
	data, err := redis.Bytes(con.Do(commandName, key, field))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return data, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// HExists determines if a hash field exists.
func (r *Redigo) HExists(_ context.Context, key, field string) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "HEXISTS"
	data, err := redis.Int64(con.Do(commandName, key, field))
	if err != nil {
		return false, errorx.E(err, errorx.CodeGateway)
	}

	return data == 1, nil
}

// HGetAll gets all the fields and values in a hash.
func (r *Redigo) HGetAll(
	_ context.Context,
	key string,
) (map[string]string, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "HGETALL"
	data, err := redis.StringMap(con.Do(commandName, key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return nil, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// HSet sets the string value of a hash field.
func (r *Redigo) HSet(
	_ context.Context,
	key, field, value string,
) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "HSET"
	res, err := redis.Int64(con.Do(commandName, key, field, value))
	if err != nil {
		return false, errorx.E(err, errorx.CodeGateway)
	}

	// 1 if field is a new field in the hash and value was set.
	// 0 if field already exists in the hash and the value was updated.
	if res != 1 && res != 0 {
		return false, nil
	}

	return true, nil
}

// HKeys gets all the fields in a hash.
func (r *Redigo) HKeys(_ context.Context, key string) ([]string, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "HKEYS"
	data, err := redis.Strings(con.Do(commandName, key))
	if err != nil {
		return nil, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// HDel deletes a hash field.
func (r *Redigo) HDel(_ context.Context, key string, fields ...string) (int64, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	params := make([]any, 0, len(fields)+1)

	params = append(params, key)
	for _, f := range fields {
		params = append(params, f)
	}

	const commandName = "HDEL"
	data, err := redis.Int64(con.Do(commandName, params...))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// Del deletes a key.
func (r *Redigo) Del(_ context.Context, key ...any) (int64, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "DEL"
	data, err := redis.Int64(con.Do(commandName, key...))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return data, nil
}

// IncrByEx increments redis key by adding expired.
func (r *Redigo) IncrByEx(
	_ context.Context,
	key string,
	by int64,
	expires int64,
) (int64, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	script := `
		local result = redis.call("INCRBY", KEYS[1], ARGV[1])
		if result then
			redis.call("EXPIRE", KEYS[1], ARGV[2])
			return result
		end
		return result
	`

	const commandName = "EVAL"
	result, err := redis.Int64(con.Do(commandName, script, 1, key, by, expires))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return result, nil
}

// LRange gets array range that we set using LPush between index Start and Stop.
func (r *Redigo) LRange(
	_ context.Context,
	key string,
	start, stop int64,
) ([][]byte, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "LRANGE"
	data, err := redis.ByteSlices(con.Do(commandName, key, start, stop))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return nil, errorx.E(err, errorx.CodeGateway)
	}

	if errors.Is(err, redis.ErrNil) {
		return nil, errorx.E("redis operation ended unsuccessfully", errorx.CodeGateway)
	}

	return data, nil
}

// LTrim array value that we set using LPush between index start and stop.
func (r *Redigo) LTrim(_ context.Context, key string, start, stop int64) error {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "LTRIM"
	data, err := redis.String(con.Do(commandName, key, start, stop))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return errorx.E(err, errorx.CodeGateway)
	}

	if errors.Is(err, redis.ErrNil) || !strings.EqualFold("OK", data) {
		return errorx.E("redis operation ended unsuccessfully", errorx.CodeGateway)
	}

	return nil
}

// SAdd add the specified members to the set stored at key.
// It returns false if key and value combination exists.
func (r *Redigo) SAdd(_ context.Context, key string, value ...string) (bool, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	args := []any{key}
	for _, v := range value {
		args = append(args, v)
	}

	const commandName = "SADD"
	data, err := redis.Int64(con.Do(commandName, args...))
	if err != nil {
		return data > 0, errorx.E(err, errorx.CodeGateway)
	}

	// data > 0 means successfully add some members.
	// No need to check redis.ErrNil because return is never (nil), but always integer.
	return data > 0, nil
}

// Publish sends message to a topic and returns numbers of subscriber that receives the message.
func (r *Redigo) Publish(_ context.Context, topic, message string) (int, error) {
	con := r.client.Get()
	defer func() {
		_ = con.Close()
	}()

	const commandName = "PUBLISH"
	res, err := redis.Int(con.Do(commandName, topic, message))
	if err != nil {
		return 0, errorx.E(err, errorx.CodeGateway)
	}

	return res, nil
}

// Close closes the client, releasing any open resources.
func (r *Redigo) Close() {
	_ = r.client.Close()
}
