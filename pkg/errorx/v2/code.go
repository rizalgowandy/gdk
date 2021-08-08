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
	CodeCircuitBreaker Code = "circuit_breaker" // Circuit breaker error.
	CodeMarshal        Code = "marshal"         // JSON marshal error.
	CodeUnmarshal      Code = "unmarshal"       // JSON unmarshal error.
	CodeConversion     Code = "conversion"      // Conversion error, e.g. string to time conversion.
	CodeEncryption     Code = "encryption"      // Encryption error.
	CodeDecryption     Code = "decryption"      // Decryption error.
	CodeDB             Code = "db"              // Database operation error.
	CodeDBScan         Code = "db_scan"         // Database scan error.
	CodeDBExec         Code = "db_exec"         // Database exec error.
	CodeDBQuery        Code = "db_query"        // Database query error.
	CodeDBBegin        Code = "db_begin"        // Database begin transaction error.
	CodeDBCommit       Code = "db_commit"       // Database commit error.
	CodeDBRollback     Code = "db_rollback"     // Database rollback error.
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
