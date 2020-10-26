package errorx

var DefaultMessage = "An internal error has occurred. Please contact technical support."

// Message returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func Message(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return Message(e.Err)
	}
	return DefaultMessage
}
