package errorx

// DefaultMessage is the string used as default response for GetMessage.
var DefaultMessage = "An internal error has occurred. Please contact technical support."

type Message string

// GetMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func GetMessage(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Message != "" {
		return string(e.Message)
	} else if ok && e.Err != nil {
		return GetMessage(e.Err)
	}

	return DefaultMessage
}
