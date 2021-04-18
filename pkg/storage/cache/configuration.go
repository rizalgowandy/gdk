package cache

// RedisConfiguration configuration.
type RedisConfiguration struct {
	// IdleConnectionLimit describes maximum idle connection before it closed.
	IdleConnectionLimit int64
	// OpenConnectionLimit describes maximum open connection can be opened.
	OpenConnectionLimit int64
	// WaitOpenConnection describes if operation should wait
	// for open connection to do operation.
	WaitOpenConnection bool
	// Addresses list of available addresses for cluster connection.
	Addresses []string
	// Enables read-only commands on slave nodes.
	ReadOnly bool
	// Allows routing read-only commands to the closest master or slave node.
	// It automatically enables ReadOnly.
	RouteByLatency bool
	// Allows routing read-only commands to the random master or slave node.
	// It automatically enables ReadOnly.
	RouteRandomly bool
	// MaxRetries describes the maximum retry if command fails.
	// Too much retry will prevent other command from being executed.
	MaxRetries int64
	// DialTimeout is in seconds.
	DialTimeout int64
	// MaxConnAge is in seconds.
	MaxConnAge int64
	// IdleTimeout is in seconds.
	IdleTimeout int64
	// MinIdleConns describes the minimum open idle connection to be keep at all time.
	MinIdleConns int64
}

// RistrettoConfiguration configuration.
type RistrettoConfiguration struct {
	// NumCounters determines the number of counters (keys) to keep that hold
	// access frequency information. It's generally a good idea to have more
	// counters than the max cache capacity, as this will improve eviction
	// accuracy and subsequent hit ratios.
	//
	// For example, if you expect your cache to hold 1,000,000 items when full,
	// NumCounters should be 10,000,000 (10x). Each counter takes up 4 bits, so
	// keeping 10,000,000 counters would require 5MB of memory.
	NumCounters int64
	// MaxCost can be considered as the cache capacity, in whatever units you
	// choose to use.
	//
	// For example, if you want the cache to have a max capacity of 100MB, you
	// would set MaxCost to 100,000,000 and pass an item's number of bytes as
	// the `cost` parameter for calls to Set. If new items are accepted, the
	// eviction process will take care of making room for the new item and not
	// overflowing the MaxCost value.
	MaxCost int64
	// BufferItems determines the size of Get buffers.
	//
	// Unless you have a rare use case, using `64` as the BufferItems value
	// results in good performance.
	BufferItems int64
	// Metrics determines whether cache statistics are kept during the cache's
	// lifetime. There *is* some overhead to keeping statistics, so you should
	// only set this flag to true when testing or throughput performance isn't a
	// major factor.
	Metrics bool
}
