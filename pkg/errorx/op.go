package errorx

// Op describes an operation, usually as the package and method,
// such as "userService.FindUserByID".
type Op string

// GetOps returns the "stack" of operations
// for each generated error.
func GetOps(err error) []Op {
	e, ok := err.(*Error)
	if !ok {
		return nil
	}

	res := []Op{
		e.Op,
	}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, GetOps(subErr)...)

	return res
}
