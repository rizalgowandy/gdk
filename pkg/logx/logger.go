package logx

type Logger interface {
	Trace(
		requestID string,
		fields map[string]any,
		message string,
	)
	Debug(
		requestID string,
		fields map[string]any,
		message string,
	)
	Info(
		requestID string,
		fields map[string]any,
		message string,
	)
	Warn(
		requestID string,
		err error,
		fields map[string]any,
		message string,
	)
	Error(
		requestID string,
		err error,
		fields map[string]any,
		message string,
	)
	Fatal(
		requestID string,
		err error,
		fields map[string]any,
		message string,
	)
}
