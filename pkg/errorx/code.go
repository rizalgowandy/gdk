package errorx

// Application error codes.
const (
	EConflict = "conflict"  // action cannot be performed
	EInternal = "internal"  // internal error
	EInvalid  = "invalid"   // validation failed
	ENotFound = "not_found" // entity does not exist
	EGateway  = "gateway"   // gateway return error
)

// Code returns the code of the root error, if available. Otherwise returns EInternal.
func Code(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return Code(e.Err)
	}
	return EInternal
}
