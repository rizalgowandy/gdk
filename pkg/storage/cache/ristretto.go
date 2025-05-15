package cache

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/syncx"
)

var (
	onceNewRistretto    syncx.Once
	onceNewRistrettoRes *Ristretto
	onceNewRistrettoErr error
	onceRistrettoClose  syncx.Once
)

// NewRistretto return a redis client.
func NewRistretto(config *RistrettoConfiguration) (*Ristretto, error) {
	onceNewRistretto.Do(func() {
		cache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: config.NumCounters,
			MaxCost:     config.MaxCost,
			BufferItems: config.BufferItems,
			Metrics:     config.Metrics,
		})
		if err != nil {
			onceNewRedigoErr = errorx.E(err)
			return
		}

		onceNewRistrettoRes = &Ristretto{
			cache: cache,
		}
	})

	return onceNewRistrettoRes, onceNewRistrettoErr
}

// Ristretto returns a in-process storage client using ristretto library.
type Ristretto struct {
	cache *ristretto.Cache
}

// Get returns the value (if any) and a boolean representing whether the value was found or not.
func (r *Ristretto) Get(_ context.Context, key string) (res any, exists bool) {
	res, exist := r.cache.Get(key)
	if !exist || res == nil {
		return nil, false
	}

	return res, true
}

// Set attempts to add the key-value item to the cache.
// If it returns false, then the Set was dropped and the key-value item isn't added to the cache.
// If it returns true, there's still a chance it could be dropped by the policy
// if its determined that the key-value item isn't worth keeping,
// but otherwise the item will be added and other items will be evicted in order to make room.
func (r *Ristretto) Set(_ context.Context, key string, value any) bool {
	return r.cache.Set(key, value, 1)
}

// SetEX works like Set but adds a key-value pair to the cache
// that will expire after the specified TTL (time to live) has passed.
// A zero value means the value never expires, which is identical to calling Set.
// A negative value is a no-op and the value is discarded.
func (r *Ristretto) SetEX(
	_ context.Context,
	key string,
	value any,
	ttl time.Duration,
) bool {
	return r.cache.SetWithTTL(key, value, 1, ttl)
}

// Del deletes the key-value item from the cache if it exists.
func (r *Ristretto) Del(_ context.Context, key string) {
	r.cache.Del(key)
}

// Clear empties the cache store and zeroes all policy counters.
func (r *Ristretto) Clear(_ context.Context) {
	r.cache.Clear()
}

// Close stops all goroutines and closes all channels.
func (r *Ristretto) Close() {
	onceRistrettoClose.Do(func() {
		r.cache.Clear()
		r.cache.Close()
	})
}
