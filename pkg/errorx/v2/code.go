package errorx

// Code defines the kind of error this is, mostly for use by systems
// that must act differently depending on the error.
type Code string

// Application error codes.
const (
	CodeUnknown    Code = ""           // Unclassified or unknown error.
	CodePermission Code = "permission" // Permission denied.
	CodeInternal   Code = "internal"   // Internal error or inconsistency.
	CodeConflict   Code = "conflict"   // Action cannot be performed.
	CodeInvalid    Code = "invalid"    // Validation failed.
	CodeNotFound   Code = "not_found"  // Entity does not exist.
	CodeGateway    Code = "gateway"    // Gateway or third party service return error.
)

// GetCode returns the code of the root error, if available. Otherwise returns CodeInternal.
func GetCode(err error) Code {
	if err == nil {
		return CodeUnknown
	}

	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return GetCode(e.Err)
	}

	return CodeInternal
}
