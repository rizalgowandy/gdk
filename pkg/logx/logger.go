package logx

type Logger interface {
	Trace(
		requestID string,
		fields map[string]interface{},
		message string,
	)
	Debug(
		requestID string,
		fields map[string]interface{},
		message string,
	)
	Info(
		requestID string,
		fields map[string]interface{},
		message string,
	)
	Warn(
		requestID string,
		err error,
		fields map[string]interface{},
		message string,
	)
	Error(
		requestID string,
		err error,
		fields map[string]interface{},
		message string,
	)
	Fatal(
		requestID string,
		err error,
		fields map[string]interface{},
		message string,
	)
}
