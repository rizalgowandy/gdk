package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/syncx"
	"github.com/rizalgowandy/gdk/pkg/tags"
)

var (
	onceNewGoRedisCluster    syncx.Once
	onceNewGoRedisClusterRes *GoRedisCluster
)

// NewGoRedisCluster returns a redis cluster client.
func NewGoRedisCluster(config *RedisConfiguration) (*GoRedisCluster, error) {
	onceNewGoRedisCluster.Do(func() {
		// Create connection to the cluster.
		// The unfilled configuration means, it will use the default configuration.
		// Tweaking configuration may increase or decrease the performance.
		rdb := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:              config.Addresses,
			NewClient:          nil,
			MaxRedirects:       0,
			ReadOnly:           config.ReadOnly,
			RouteByLatency:     config.RouteByLatency,
			RouteRandomly:      config.RouteRandomly,
			ClusterSlots:       nil,
			Dialer:             nil,
			OnConnect:          nil,
			Username:           "",
			Password:           "",
			MaxRetries:         int(config.MaxRetries),
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        time.Duration(config.DialTimeout) * time.Second,
			ReadTimeout:        0,
			WriteTimeout:       0,
			PoolSize:           int(config.OpenConnectionLimit),
			MinIdleConns:       int(config.MinIdleConnection),
			MaxConnAge:         time.Duration(config.MaxConnAge) * time.Second,
			PoolTimeout:        0,
			IdleTimeout:        time.Duration(config.IdleTimeout) * time.Second,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
		})

		onceNewGoRedisClusterRes = &GoRedisCluster{
			client: rdb,
		}
	})

	return onceNewGoRedisClusterRes, nil
}

// GoRedisCluster returns a redis cluster client using go-redis library.
type GoRedisCluster struct {
	client *redis.ClusterClient
}

// Get gets the value from redis in []byte form.
func (r *GoRedisCluster) Get(ctx context.Context, key string) ([]byte, error) {
	const op errorx.Op = "cache/GoRedisCluster.Get"

	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return []byte(res), nil
}

// SetEX sets the value to a key with timeout in seconds.
func (r *GoRedisCluster) SetEX(
	ctx context.Context,
	key string,
	seconds int64,
	value string,
) error {
	const op errorx.Op = "cache/GoRedisCluster.SetEX"

	_, err := r.client.SetEX(ctx, key, value, time.Duration(seconds)*time.Second).Result()
	if err != nil {
		return errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key}, string(op)+" success")
	return nil
}

// SetNX sets a value to a key with specified timeouts.
// SetNX returns false if the key exists.
func (r *GoRedisCluster) SetNX(
	ctx context.Context,
	key string,
	seconds int64,
	value string,
) (bool, error) {
	const op errorx.Op = "cache/GoRedisCluster.SetNX"

	res, err := r.client.SetNX(ctx, key, value, time.Duration(seconds)*time.Second).Result()
	if err != nil {
		return false, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return res, nil
}

// Exists checks whether the key exists in redis.
func (r *GoRedisCluster) Exists(ctx context.Context, key string) (bool, error) {
	const op errorx.Op = "cache/GoRedisCluster.Exists"

	res, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return res > 0, nil
}

// Expire sets the ttl of a key to specified value in seconds.
func (r *GoRedisCluster) Expire(
	ctx context.Context,
	key string,
	seconds int64,
) (bool, error) {
	const op errorx.Op = "cache/GoRedisCluster.Expire"

	res, err := r.client.Expire(ctx, key, time.Duration(seconds)*time.Second).Result()
	if err != nil {
		return false, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return res, nil
}

// TTL gets the time to live of a key / expiry time.
func (r *GoRedisCluster) TTL(ctx context.Context, key string) (int64, error) {
	const op errorx.Op = "cache/GoRedisCluster.TTL"

	res, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		return 0, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return int64(res), nil
}

// HGet gets the value of a hash field.
func (r *GoRedisCluster) HGet(ctx context.Context, key, field string) ([]byte, error) {
	const op errorx.Op = "cache/GoRedisCluster.HGet"

	res, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		return nil, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return []byte(res), nil
}

// HExists determines if a hash field exists.
func (r *GoRedisCluster) HExists(ctx context.Context, key, field string) (bool, error) {
	const op errorx.Op = "cache/GoRedisCluster.HExists"

	res, err := r.client.HExists(ctx, key, field).Result()
	if err != nil {
		return false, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return res, nil
}

// HSet sets the string value of a hash field.
func (r *GoRedisCluster) HSet(ctx context.Context, key, field, value string) (bool, error) {
	const op errorx.Op = "cache/GoRedisCluster.HSet"

	res, err := r.client.HSet(ctx, key, field, value).Result()
	if err != nil {
		return false, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return res > 0, nil
}

// Del deletes a key.
func (r *GoRedisCluster) Del(ctx context.Context, key ...interface{}) (int64, error) {
	const op errorx.Op = "cache/GoRedisCluster.Del"

	stdKeys := make([]string, len(key))
	for i, v := range key {
		stdKey, ok := v.(string)
		if ok {
			stdKeys[i] = stdKey
		}
	}

	res, err := r.client.Del(ctx, stdKeys...).Result()
	if err != nil {
		return 0, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{
		tags.Key:    key,
		tags.Detail: res,
	}, string(op)+" success")
	return res, nil
}

// Close closes the client, releasing any open resources.
func (r *GoRedisCluster) Close() {
	_ = r.client.Close()
}
