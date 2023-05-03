package errorx

import (
	"fmt"
)

// E builds an error value from its arguments.
// There must be at least one argument or E panics.
// The type of each argument determines its meaning.
// If more than one argument of a given type is presented,
// only the last one is recorded.
//
// The types are:
//
//	 errorx.Op
//		    The operation being performed, usually the method
//			being invoked (Get, Put, etc.).
//		string
//			Treated as an error message and assigned to the Message.
//		errorx.Code
//			The class of error, such as permission failure.
//		error
//			The underlying error that triggered this one.
//
// If the error is printed, only those items that have been
// set to non-zero values will appear in the result.
//
// If Code is not specified or Unknown, we set it to the Code of
// the underlying error.
//
// If Message is not filled, we set it to the Message of
// the underlying error.
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errorx.E with no arguments")
	}

	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			e.Message = arg
		case Op:
			e.Op = arg
		case Code:
			e.Code = arg
		case *Error:
			// Make a copy
			errorCopy := *arg
			e.Err = &errorCopy
		case error:
			e.Err = arg
		default:
			panic(fmt.Sprintf("unknown type %T, value %v in error call", arg, arg))
		}
	}

	prev, ok := e.Err.(*Error)
	if !ok {
		return e
	}
	// The previous error was also one of ours.
	// Suppress duplications so the message won't
	// contain the same code and message twice.
	if prev.Code == e.Code {
		prev.Code = Unknown
	}
	// If this error has Code unset or Unknown, pull up the inner one.
	if e.Code == Unknown {
		e.Code = prev.Code
		prev.Code = Unknown
	}
	// If this error has empty Message, pull up the inner one.
	if e.Message == "" {
		e.Message = prev.Message
		prev.Message = ""
	}

	return e
}

// Match compares its two error arguments.
// It can be used to check for expected errors in tests.
// Both arguments must have underlying
// type *Error or Match will return false.
// Otherwise it returns true if
// every non-zero element of the first error is equal to
// the corresponding element of the second.
// If the Err field is a *Error, Match recurs on that field;
// otherwise it compares the strings returned by the Error methods.
// Elements that are in the second argument but not present in
// the first are ignored.
//
// Example:
//
//	Match(errors.E(errorx.Permission, "message"), err)
//	tests whether err is an Error with Code=Permission and Message=message.
func Match(err1, err2 error) bool {
	e1, ok := err1.(*Error)
	if !ok {
		return false
	}
	e2, ok := err2.(*Error)
	if !ok {
		return false
	}

	// Compare properties.
	if isMessageUnequal(e1.Message, e2.Message) ||
		isOpUnequal(e1.Op, e2.Op) ||
		isCodeUnequal(e1.Code, e2.Code) {
		return false
	}

	if e1.Err != nil {
		if _, ok := e1.Err.(*Error); ok {
			return Match(e1.Err, e2.Err)
		}
		if e2.Err == nil || e1.Err.Error() != e2.Err.Error() {
			return false
		}
	}

	return true
}

// Is reports whether err is an *Error of the given Code.
// If err is nil then Is returns false.
func Is(code Code, err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Code != Unknown {
		return e.Code == code
	}

	if e.Err != nil {
		return Is(code, e.Err)
	}

	return false
}

func isMessageUnequal(msg1, msg2 string) bool {
	return msg1 != "" && msg1 != msg2
}

func isOpUnequal(op1, op2 Op) bool {
	return op1 != "" && op1 != op2
}

func isCodeUnequal(code1, code2 Code) bool {
	return code1 != Unknown && code1 != code2
}
