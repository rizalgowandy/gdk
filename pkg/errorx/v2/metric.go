package errorx

// MetricStatus defines the kind of error should be tracked or not.
// Useful for alerting system.
type MetricStatus string

const (
	MetricSuccess       MetricStatus = "success"
	MetricError         MetricStatus = "error"
	MetricExpectedError MetricStatus = "expected_error"
)
