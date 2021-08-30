package errorx

import (
	"errors"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/rizalgowandy/gdk/pkg/fn"
)

const callerSkip = 2

// E for creating new error.
// error should always be the first param.
func E(args ...interface{}) error {
	if len(args) == 0 {
		return Errorf("errorx.E: bad call without args from file=%s", fn.Line(callerSkip))
	}

	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case *Error:
			// Copy and put the errors back.
			errCopy := *arg
			e = &errCopy
			e.OpTraces = append([]Op{Op(fn.Name(callerSkip))}, e.OpTraces...)

		case error:
			e.Err = arg
			e.Line = Line(fn.Line(callerSkip))
			e.OpTraces = append([]Op{Op(fn.Name(callerSkip))}, e.OpTraces...)

		case string:
			e.Err = Errorf(arg)
			e.Line = Line(fn.Line(callerSkip))
			e.OpTraces = append([]Op{Op(fn.Name(callerSkip))}, e.OpTraces...)

		case Code:
			// New code will always replace the old code.
			e.Code = arg

		case Fields:
			// Previous fields is empty.
			// Replace with arg directly.
			if e.Fields == nil {
				e.Fields = arg
				continue
			}

			// Merge fields.
			// If there is duplicate fields,
			// The e.Fields has higher priority and won't be replaced with arg.
			if err := mergo.Merge(&e.Fields, arg); err != nil {
				e.Fields = arg
			}

		case Op:
			// For backward compatibility.
			// Client is no longer to pass Op manually as an argument but will be filled automatically.

		case Message:
			e.Message = arg

		case MetricStatus:
			e.MetricStatus = arg

		default:
			// The default error is unknown.
			msg := fmt.Sprintf("errorx.E: bad call from file=%s args=%v", fn.Line(callerSkip), args)
			return Errorf(msg+"; unknown_type=%T value=%v", arg, arg)
		}
	}
	return e
}

// Match compares its two error arguments.
// It can be used to check for expected errors in tests.
// Both arguments must have underlying type *Error or Match will return false.
// Otherwise it returns true if every non-zero element of the first error is equal to
// the corresponding element of the second.
// If the Err field is a *Error, Match recurs on that field;
// otherwise it compares the strings returned by the Error methods.
// Elements that are in the second argument but not present in the first are ignored.
func Match(errs1, errs2 error) bool {
	if errs1 == nil && errs2 == nil {
		return true
	}

	if errs1 != nil {
		err1, ok := errs1.(*Error)
		if ok {
			errs1 = err1.Err
		}
	} else {
		errs1 = errors.New("nil")
	}

	if errs2 != nil {
		err2, ok := errs2.(*Error)
		if ok {
			errs2 = err2.Err
		}
	} else {
		errs2 = errors.New("nil")
	}

	if errs1.Error() != errs2.Error() {
		return false
	}
	return true
}

// Is reports whether err is an *Error of the given Code.
// If err is nil then Is returns false.
func Is(err error, code Code) bool {
	if err == nil {
		return false
	}

	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Code != CodeUnknown {
		return e.Code == code
	}

	if e.Err != nil {
		return Is(e.Err, code)
	}

	return false
}
