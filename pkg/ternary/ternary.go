package ternary

// Str returns string based on condition.
func Str(condition bool, yes, no string) string {
	if condition {
		return yes
	}
	return no
}

// Int64 returns int64 based on condition.
func Int64(condition bool, yes, no int64) int64 {
	if condition {
		return yes
	}
	return no
}

// Int returns int based on condition.
func Int(condition bool, yes, no int) int {
	if condition {
		return yes
	}
	return no
}

func Float64(condition bool, yes, no float64) float64 {
	if condition {
		return yes
	}
	return no
}
