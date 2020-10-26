package errorx

// Code defines the kind of error this is, mostly for use by systems
// that must act differently depending on the error.
type Code string

// Application error codes.
const (
	Unknown    Code = ""           // Unclassified or unknown error. This value is not printed in the error message.
	Permission Code = "permission" // Permission denied.
	Internal   Code = "internal"   // Internal error or inconsistency.
	Conflict   Code = "conflict"   // Action cannot be performed.
	Invalid    Code = "invalid"    // Validation failed.
	NotFound   Code = "not_found"  // Entity does not exist.
	Gateway    Code = "gateway"    // Gateway or third party service return error.
)

// GetCode returns the code of the root error, if available. Otherwise returns Internal.
func GetCode(err error) Code {
	if err == nil {
		return Unknown
	}

	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return GetCode(e.Err)
	}

	return Internal
}
