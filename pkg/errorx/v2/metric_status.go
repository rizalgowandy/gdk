package errorx

// MetricStatus defines the kind of error should be tracked or not.
// Useful for alerting system.
type MetricStatus string

const (
	MetricStatusSuccess     MetricStatus = "success"
	MetricStatusErr         MetricStatus = "error"
	MetricStatusExpectedErr MetricStatus = "expected_error"
)
