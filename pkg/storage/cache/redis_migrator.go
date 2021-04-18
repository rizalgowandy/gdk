package cache

import (
	"context"

	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/resync"
	"github.com/peractio/gdk/pkg/tags"
)

var (
	onceNewRedisMigratorClient    resync.Once
	onceNewRedisMigratorClientRes *RedisMigratorClient
	onceNewRedisMigratorClientErr error
)

// RedisMigratorClient is a redis client to migrate from one instance to another.
// All the create commands will be executed to the new instance.
// All the update and delete commands will be executed on both instances.
// While the read commands will read from the new instance first,
// and if the result is not found, it will attempt to read from the old one.
type RedisMigratorClient struct {
	oldClient RedisClientItf
	newClient RedisClientItf
}

// NewRedisMigratorClient return a redis migrator client.
func NewRedisMigratorClient(config *RedisConfiguration) (*RedisMigratorClient, error) {
	onceNewRedisMigratorClient.Do(func() {
		const op errorx.Op = "cache.NewRedisMigratorClient"

		// Create the old client.
		oldClient, err := NewRedigoClient(config)
		if err != nil {
			onceNewRedisMigratorClientErr = errorx.E(err, op)
			return
		}

		// Create the new client.
		newClient, err := NewRedisClusterClient(config)
		if err != nil {
			onceNewRedisMigratorClientErr = errorx.E(err, op)
			return
		}

		onceNewRedisMigratorClientRes = &RedisMigratorClient{
			oldClient: oldClient,
			newClient: newClient,
		}
	})

	return onceNewRedisMigratorClientRes, onceNewRedisMigratorClientErr
}

// Get gets the value from redis in []byte form.
func (r *RedisMigratorClient) Get(ctx context.Context, key string) ([]byte, error) {
	const op errorx.Op = "cache/RedisMigratorClient.Get"

	// Read from new client.
	res, err := r.newClient.Get(ctx, key)
	if err == nil && res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
		return res, nil
	}

	// Read from old client.
	res, err = r.oldClient.Get(ctx, key)
	if err != nil {
		return nil, errorx.E(err, op)
	}
	if res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}
	return res, nil
}

// SetEX sets the value to a key with timeout in seconds.
func (r *RedisMigratorClient) SetEX(
	ctx context.Context,
	key string,
	seconds int64,
	value string,
) error {
	const op errorx.Op = "cache/RedisMigratorClient.SetEX"

	// Create to new client.
	err := r.newClient.SetEX(ctx, key, seconds, value)
	if err != nil {
		return errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
	return nil
}

// Exists checks whether the key exists in redis.
func (r *RedisMigratorClient) Exists(ctx context.Context, key string) (bool, error) {
	const op errorx.Op = "cache/RedisMigratorClient.Exists"

	// Read from new client.
	res, err := r.newClient.Exists(ctx, key)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
		return true, nil
	}

	// Read from old client.
	res, err = r.oldClient.Exists(ctx, key)
	if err != nil {
		return false, errorx.E(err, op)
	}
	if res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}
	return res, nil
}

// Expire sets the ttl of a key to specified value in seconds.
func (r *RedisMigratorClient) Expire(ctx context.Context, key string, seconds int64) (bool, error) {
	const op errorx.Op = "cache/RedisMigratorClient.Expire"

	// Update to old client.
	res, err := r.oldClient.Expire(ctx, key, seconds)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}

	// Update to new client.
	res, err = r.newClient.Expire(ctx, key, seconds)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
	}

	return res, err
}

// TTL gets the time to live of a key / expiry time.
func (r *RedisMigratorClient) TTL(ctx context.Context, key string) (int64, error) {
	const op errorx.Op = "cache/RedisMigratorClient.TTL"

	// Read from new client.
	res, err := r.newClient.TTL(ctx, key)
	if err == nil && res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
		return res, nil
	}

	// Read from old client.
	res, err = r.oldClient.TTL(ctx, key)
	if err != nil {
		return 0, errorx.E(err, op)
	}
	if res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}
	return res, nil
}

// HGet gets the value of a hash field.
func (r *RedisMigratorClient) HGet(ctx context.Context, key, field string) ([]byte, error) {
	const op errorx.Op = "cache/RedisMigratorClient.HGet"

	// Read from new client.
	res, err := r.newClient.HGet(ctx, key, field)
	if err == nil && res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
		return res, nil
	}

	// Read from old client.
	res, err = r.oldClient.HGet(ctx, key, field)
	if err != nil {
		return nil, errorx.E(err, op)
	}
	if res != nil {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}
	return res, nil
}

// HExists determines if a hash field exists.
func (r *RedisMigratorClient) HExists(ctx context.Context, key, field string) (bool, error) {
	const op errorx.Op = "cache/RedisMigratorClient.HGet"

	// Read from new client.
	res, err := r.newClient.HExists(ctx, key, field)
	if err == nil && res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
		return true, nil
	}

	// Read from old client.
	res, err = r.oldClient.HExists(ctx, key, field)
	if err != nil {
		return false, errorx.E(err, op)
	}
	if res {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}
	return res, nil
}

// HSet sets the string value of a hash field.
func (r *RedisMigratorClient) HSet(ctx context.Context, key, field, value string) (bool, error) {
	const op errorx.Op = "cache/RedisMigratorClient.HSet"

	// Create to new client.
	res, err := r.newClient.HSet(ctx, key, field, value)
	if err != nil {
		return false, errorx.E(err, op)
	}

	logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
	return res, nil
}

// Del deletes a key.
func (r *RedisMigratorClient) Del(ctx context.Context, key ...interface{}) (int64, error) {
	const op errorx.Op = "cache/RedisMigratorClient.Del"

	// Update to old client.
	res, err := r.oldClient.Del(ctx, key...)
	if err == nil && res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "old"}, string(op)+" success")
	}

	// Update to new client.
	res, err = r.newClient.Del(ctx, key...)
	if err == nil && res > 0 {
		logx.DBG(ctx, logx.KV{tags.Key: key, tags.Client: "new"}, string(op)+" success")
	}

	return res, err
}
