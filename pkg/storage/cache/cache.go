package cache

import (
	"context"
	"time"
)

//go:generate mockgen -destination=cache_mock.go -package=cache -source=cache.go

// RedisClientItf is a client to interact with Redis.
type RedisClientItf interface {
	// Get gets the value from redis in []byte form.
	Get(ctx context.Context, key string) ([]byte, error)
	// SetEX sets the value to a key with timeout in seconds.
	SetEX(ctx context.Context, key string, seconds int64, value string) error
	// SetNX sets a value to a key with specified timeouts.
	// SetNX returns false if the key exists.
	SetNX(ctx context.Context, key string, seconds int64, value string) (bool, error)
	// Exists checks whether the key exists in redis.
	Exists(ctx context.Context, key string) (bool, error)
	// Expire sets the TTL of a key to specified value in seconds.
	Expire(ctx context.Context, key string, seconds int64) (bool, error)
	// TTL gets the time to live of a key / expiry time.
	TTL(ctx context.Context, key string) (int64, error)
	// HGet gets the value of a hash field.
	HGet(ctx context.Context, key, field string) ([]byte, error)
	// HExists determines if a hash field exists.
	HExists(ctx context.Context, key, field string) (bool, error)
	// HSet sets the string value of a hash field.
	HSet(ctx context.Context, key, field, value string) (bool, error)
	// Del deletes a key.
	Del(ctx context.Context, key ...any) (int64, error)
	// Close closes the client, releasing any open resources.
	Close()
}

// RistrettoClientItf is a in-process or local cache storage client.
type RistrettoClientItf interface {
	// Get returns the value (if any) and a boolean representing whether the value was found or not.
	Get(ctx context.Context, key string) (res any, exists bool)
	// Set attempts to add the key-value item to the cache.
	// If it returns false, then the Set was dropped and the key-value item isn't added to the cache.
	// If it returns true, there's still a chance it could be dropped by the policy
	// if its determined that the key-value item isn't worth keeping,
	// but otherwise the item will be added and other items will be evicted in order to make room.
	Set(ctx context.Context, key string, value any) bool
	// SetEX works like Set but adds a key-value pair to the cache
	// that will expire after the specified TTL (time to live) has passed.
	// A zero value means the value never expires, which is identical to calling Set.
	// A negative value is a no-op and the value is discarded.
	SetEX(ctx context.Context, key string, value any, TTL time.Duration) bool
	// Del deletes the key-value item from the cache if it exists.
	Del(ctx context.Context, key string)
	// Clear empties the cache store and zeroes all policy counters.
	Clear(ctx context.Context)
	// Close stops all goroutines and closes all channels.
	Close()
}
