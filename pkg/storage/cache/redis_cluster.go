package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/resync"
	"github.com/peractio/gdk/pkg/tags"
)

var (
	onceNewRedisClusterClient    resync.Once
	onceNewRedisClusterClientRes *RedisClusterClient
)

// RedisClusterClient returns a redis cluster client using go-redis library.
type RedisClusterClient struct {
	client *redis.ClusterClient
}

// NewRedisClusterClient returns a redis cluster client.
func NewRedisClusterClient(config *RedisConfiguration) (*RedisClusterClient, error) {
	onceNewRedisClusterClient.Do(func() {
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
			MinIdleConns:       int(config.MinIdleConns),
			MaxConnAge:         time.Duration(config.MaxConnAge) * time.Second,
			PoolTimeout:        0,
			IdleTimeout:        time.Duration(config.IdleTimeout) * time.Second,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
		})

		onceNewRedisClusterClientRes = &RedisClusterClient{
			client: rdb,
		}
	})

	return onceNewRedisClusterClientRes, nil
}

// Get gets the value from redis in []byte form.
func (r *RedisClusterClient) Get(ctx context.Context, key string) ([]byte, error) {
	const op errorx.Op = "cache/RedisClusterClient.Get"

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
func (r *RedisClusterClient) SetEX(
	ctx context.Context,
	key string,
	seconds int64,
	value string,
) error {
	const op errorx.Op = "cache/RedisClusterClient.SetEX"

	_, err := r.client.SetEX(ctx, key, value, time.Duration(seconds)*time.Second).Result()
	if err != nil {
		return errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key}, string(op)+" success")
	return nil
}

// Exists checks whether the key exists in redis.
func (r *RedisClusterClient) Exists(ctx context.Context, key string) (bool, error) {
	const op errorx.Op = "cache/RedisClusterClient.Exists"

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
func (r *RedisClusterClient) Expire(ctx context.Context, key string, seconds int64) (bool, error) {
	const op errorx.Op = "cache/RedisClusterClient.Expire"

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
func (r *RedisClusterClient) TTL(ctx context.Context, key string) (int64, error) {
	const op errorx.Op = "cache/RedisClusterClient.TTL"

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
func (r *RedisClusterClient) HGet(ctx context.Context, key, field string) ([]byte, error) {
	const op errorx.Op = "cache/RedisClusterClient.HGet"

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
func (r *RedisClusterClient) HExists(ctx context.Context, key, field string) (bool, error) {
	const op errorx.Op = "cache/RedisClusterClient.HExists"

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
func (r *RedisClusterClient) HSet(ctx context.Context, key, field, value string) (bool, error) {
	const op errorx.Op = "cache/RedisClusterClient.HSet"

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
func (r *RedisClusterClient) Del(ctx context.Context, key ...interface{}) (int64, error) {
	const op errorx.Op = "cache/RedisClusterClient.Del"

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
