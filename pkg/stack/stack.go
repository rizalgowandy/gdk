package stack

import (
	"bytes"
	"strings"
)

var (
	PanicKeyword           = []byte("src/runtime/panic.go")
	CompanyKeyword         = []byte("github.com/peractio")
	FunctionPackageKeyword = []byte("/github.com/peractio/gdk/pkg/stack")
)

// Trim removes unnecessary stack trace.
// Only take stack trace with current company keyword.
// Also, excludes lines with function package keyword.
func Trim(stack []byte) []byte {
	// remove all log before panic keyword.
	idx := bytes.Index(stack, PanicKeyword)
	if idx == -1 {
		idx = 0
	}
	stack = stack[idx:]

	// remove all log before company keyword.
	idx = bytes.Index(stack, CompanyKeyword)
	if idx == -1 {
		idx = 0
	}
	stack = stack[idx:]

	// remove all log before current function location keyword.
	idx = bytes.Index(stack, FunctionPackageKeyword)
	if idx != -1 {
		newlineIdx := bytes.Index(stack[idx:], []byte("\n"))
		if newlineIdx != -1 {
			idx += newlineIdx + 1
			return stack[:idx]
		}
	}

	return stack
}

// ToArr split stack trace by newline.
// Only trim all characters before company keyword.
//
// Example:
// Input  => /home/peractio/go/src/github.com/peractio/gdk/pkg/stack.go 130
// Result => gdk/pkg/stack.go 130
//
func ToArr(stack []byte) []string {
	msg := string(stack)
	arr := strings.Split(msg, "\n")

	var res []string
	for _, v := range arr {
		if v == "" {
			continue
		}

		tmp := []byte(v)
		idx := bytes.Index(tmp, CompanyKeyword)
		if idx == -1 {
			idx = 0
		}
		tmp = tmp[idx:]
		res = append(res, string(tmp))
	}

	return res
}
