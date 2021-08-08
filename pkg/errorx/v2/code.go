package errorx

// Code defines the kind of error this is, mostly for use by systems
// that must act differently depending on the error.
type Code string

// Application error codes.
const (
	CodeUnknown        Code = ""                // Unclassified or unknown error.
	CodePermission     Code = "permission"      // Permission denied.
	CodeInternal       Code = "internal"        // Internal error or inconsistency.
	CodeConflict       Code = "conflict"        // Action cannot be performed.
	CodeInvalid        Code = "invalid"         // Validation failed.
	CodeNotFound       Code = "not_found"       // Entity does not exist.
	CodeGateway        Code = "gateway"         // Gateway or third party service return error.
	CodeConfig         Code = "config"          // Wrong configuration.
	CodeDB             Code = "db"              // Database operation error.
	CodeCircuitBreaker Code = "circuit_breaker" // Circuit breaker error.
	CodeMarshal        Code = "marshal"
	CodeUnmarshal      Code = "unmarshal"
	CodeConversion     Code = "conversion"
	CodeEncryption     Code = "encryption"
	CodeDecryption     Code = "decryption"
	CodeDBScan         Code = "db_scan"
	CodeDBExec         Code = "db_exec"
	CodeDBQuery        Code = "db_query"
	CodeDBBegin        Code = "db_begin"
	CodeDBCommit       Code = "db_commit"
	CodeDBRollback     Code = "db_rollback"
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
