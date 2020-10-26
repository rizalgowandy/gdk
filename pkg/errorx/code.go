package errorx

// Application error codes.
const (
	ECONFLICT = "conflict"  // action cannot be performed
	EINTERNAL = "internal"  // internal error
	EINVALID  = "invalid"   // validation failed
	ENOTFOUND = "not_found" // entity does not exist
)

// Code returns the code of the root error, if available. Otherwise returns EINTERNAL.
func Code(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return Code(e.Err)
	}
	return EINTERNAL
}

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
	return "An internal error has occurred. Please contact technical support."
}
