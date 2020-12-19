package errorx

// Error defines a standard application error.
type Error struct {
	// Underlying error.
	Err error

	// Codes used for Errs to identify known errors in the application
	// If the error is expected by Errs object, the errors will be shown as listed in Codes
	Code Code

	// Fields is a fields context similar to logrus.Fields
	// Can be used for adding more context to the errors
	Fields Fields

	// Op is operation of error
	Op Op

	// OpTraces is a trace of operations
	OpTraces []Op

	// Message is a human-readable message.
	Message Message

	// Line describes current error original line.
	// Only injected when the underlying error is from standard error.
	Line Line

	// MetricStatus defines the kind of error should be tracked or not.
	MetricStatus MetricStatus
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	return e.Err.Error()
}

// GetFields return available fields in errors.
func (e *Error) GetFields() Fields {
	return e.Fields
}
