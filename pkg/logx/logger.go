package logx

// Logger is an interface that provides logging functionality.
type Logger interface {
	Trace(
		requestID string,
		actorID int,
		fields map[string]any,
		message string,
	)
	Debug(
		requestID string,
		actorID int,
		fields map[string]any,
		message string,
	)
	Info(
		requestID string,
		actorID int,
		fields map[string]any,
		message string,
	)
	Warn(
		requestID string,
		actorID int,
		err error,
		fields map[string]any,
		message string,
	)
	Error(
		requestID string,
		actorID int,
		err error,
		fields map[string]any,
		message string,
	)
	Fatal(
		requestID string,
		actorID int,
		err error,
		fields map[string]any,
		message string,
	)
}
