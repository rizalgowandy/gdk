package cache

// Time use for redis ttl which by default is in time.Second.
const (
	TTLRedisOneSecond          = int64(1)
	TTLRedisOneMinute          = 60 * TTLRedisOneSecond
	TTLRedisFiveMinutes        = 5 * TTLRedisOneMinute
	TTLRedisTenMinutes         = 10 * TTLRedisOneMinute
	TTLRedisFifteenMinutes     = 15 * TTLRedisOneMinute
	TTLRedisThirtyMinutes      = 30 * TTLRedisOneMinute
	TTLRedisSeventyFiveMinutes = 75 * TTLRedisOneMinute
	TTLRedisOneHour            = 60 * TTLRedisOneMinute
	TTLRedisHalfDay            = 12 * TTLRedisOneHour
	TTLRedisOneDay             = 24 * TTLRedisOneHour
	TTLRedisThreeDays          = 3 * TTLRedisOneDay
	TTLRedisOneWeek            = 7 * TTLRedisOneDay
	TTLRedisOneMonth           = 30 * TTLRedisOneDay
)
