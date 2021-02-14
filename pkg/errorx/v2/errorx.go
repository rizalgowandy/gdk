package errorx

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

const COMPANY = "peractio"

// E for creating new error.
// error should always be the first param.
func E(args ...interface{}) error {
	if len(args) == 0 {
		_, file, line, _ := runtime.Caller(1)
		file = file[strings.Index(file, COMPANY)+len(COMPANY):]
		return Errorf("errorx.E: bad call without args from file=%s:%d", file, line)
	}

	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case *Error:
			// Copy and put the errors back.
			errCopy := *arg
			e = &errCopy

		case error:
			e.Err = arg
			_, file, line, _ := runtime.Caller(1)
			file = file[strings.Index(file, COMPANY)+len(COMPANY):]
			e.Line = Line(fmt.Sprintf("%s:%d", file, line))

		case string:
			e.Err = Errorf(arg)
			_, file, line, _ := runtime.Caller(1)
			file = file[strings.Index(file, COMPANY)+len(COMPANY):]
			e.Line = Line(fmt.Sprintf("%s:%d", file, line))

		case Code:
			// New code will always replace the old code.
			e.Code = arg

		case Fields:
			// Fields cannot be appended.
			// New fields will always replace the old fields.
			e.Fields = arg

		case Op:
			e.OpTraces = append([]Op{arg}, e.OpTraces...)

		case Message:
			e.Message = arg

		case MetricStatus:
			e.MetricStatus = arg

		default:
			// The default error is unknown.
			_, file, line, _ := runtime.Caller(1)
			file = file[strings.Index(file, COMPANY)+len(COMPANY):]
			msg := fmt.Sprintf("errorx.E: bad call from file=%s:%d args=%v", file, line, args)
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
func Is(code Code, err error) bool {
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
		return Is(code, e.Err)
	}

	return false
}
