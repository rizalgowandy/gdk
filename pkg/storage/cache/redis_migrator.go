package cache

import (
	"context"

	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/fn"
	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/syncx"
	"github.com/rizalgowandy/gdk/pkg/tags"
)

var (
	onceNewRedisMigrator    syncx.Once
	onceNewRedisMigratorRes *RedisMigrator
	onceNewRedisMigratorErr error
)

// NewRedisMigrator return a redis migrator client.
func NewRedisMigrator(origin, destination RedisClientItf) (*RedisMigrator, error) {
	onceNewRedisMigrator.Do(func() {
		onceNewRedisMigratorRes = &RedisMigrator{
			origin:      origin,
			destination: destination,
		}
	})

	return onceNewRedisMigratorRes, onceNewRedisMigratorErr
}

// RedisMigrator is a redis client to migrate from one instance to another.
// All the create commands will be executed to the new instance.
// All the update and delete commands will be executed on both instances.
// While the read commands will read from the new instance first,
// and if the result is not found, it will attempt to read from the old one.
type RedisMigrator struct {
	origin      RedisClientItf
	destination RedisClientItf
}

// Get gets the value from redis in []byte form.
func (r *RedisMigrator) Get(ctx context.Context, key string) ([]byte, error) {
	// Read from new client.
	res, err := r.destination.Get(ctx, key)
	if err == nil && res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
		return res, nil
	}

	// Read from old client.
	res, err = r.origin.Get(ctx, key)
	if err != nil {
		return nil, errorx.E(err)
	}
	if res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}
	return res, nil
}

// SetEX sets the value to a key with timeout in seconds.
func (r *RedisMigrator) SetEX(
	ctx context.Context,
	key string,
	seconds int64,
	value string,
) error {
	// Create to new client.
	err := r.destination.SetEX(ctx, key, seconds, value)
	if err != nil {
		return errorx.E(err)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
	return nil
}

// SetNX sets a value to a key with specified timeouts.
// SetNX returns false if the key exists.
func (r *RedisMigrator) SetNX(
	ctx context.Context,
	key string,
	seconds int64,
	value string,
) (bool, error) {
	// Read from old client.
	res, err := r.destination.SetNX(ctx, key, seconds, value)
	if err != nil {
		return false, errorx.E(err)
	}
	if res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
		return true, nil
	}

	// Create to new client.
	res, err = r.destination.SetNX(ctx, key, seconds, value)
	if err != nil {
		return false, errorx.E(err)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
	return res, nil
}

// Exists checks whether the key exists in redis.
func (r *RedisMigrator) Exists(ctx context.Context, key string) (bool, error) {
	// Read from new client.
	res, err := r.destination.Exists(ctx, key)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
		return true, nil
	}

	// Read from old client.
	res, err = r.origin.Exists(ctx, key)
	if err != nil {
		return false, errorx.E(err)
	}
	if res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}
	return res, nil
}

// Expire sets the ttl of a key to specified value in seconds.
func (r *RedisMigrator) Expire(ctx context.Context, key string, seconds int64) (bool, error) {
	// Update to old client.
	res, err := r.origin.Expire(ctx, key, seconds)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}

	// Update to new client.
	res, err = r.destination.Expire(ctx, key, seconds)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
	}

	return res, err
}

// TTL gets the time to live of a key / expiry time.
func (r *RedisMigrator) TTL(ctx context.Context, key string) (int64, error) {
	// Read from new client.
	res, err := r.destination.TTL(ctx, key)
	if err == nil && res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
		return res, nil
	}

	// Read from old client.
	res, err = r.origin.TTL(ctx, key)
	if err != nil {
		return 0, errorx.E(err)
	}
	if res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}
	return res, nil
}

// HGet gets the value of a hash field.
func (r *RedisMigrator) HGet(ctx context.Context, key, field string) ([]byte, error) {
	// Read from new client.
	res, err := r.destination.HGet(ctx, key, field)
	if err == nil && res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
		return res, nil
	}

	// Read from old client.
	res, err = r.origin.HGet(ctx, key, field)
	if err != nil {
		return nil, errorx.E(err)
	}
	if res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}
	return res, nil
}

// HExists determines if a hash field exists.
func (r *RedisMigrator) HExists(ctx context.Context, key, field string) (bool, error) {
	// Read from new client.
	res, err := r.destination.HExists(ctx, key, field)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
		return true, nil
	}

	// Read from old client.
	res, err = r.origin.HExists(ctx, key, field)
	if err != nil {
		return false, errorx.E(err)
	}
	if res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}
	return res, nil
}

// HSet sets the string value of a hash field.
func (r *RedisMigrator) HSet(ctx context.Context, key, field, value string) (bool, error) {
	// Create to new client.
	res, err := r.destination.HSet(ctx, key, field, value)
	if err != nil {
		return false, errorx.E(err)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
	return res, nil
}

// Del deletes a key.
func (r *RedisMigrator) Del(ctx context.Context, key ...interface{}) (int64, error) {
	// Update to old client.
	res, err := r.origin.Del(ctx, key...)
	if err == nil && res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "origin"}, fn.Name()+" success")
	}

	// Update to new client.
	res, err = r.destination.Del(ctx, key...)
	if err == nil && res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "destination"}, fn.Name()+" success")
	}

	return res, err
}

// Close closes the client, releasing any open resources.
func (r *RedisMigrator) Close() {
	r.origin.Close()
	r.destination.Close()
}
