package stack

import (
	"bytes"
	"strings"
)

var (
	panicKeyword           = []byte("src/runtime/panic.go")
	companyKeyword         = []byte("github.com/rizalgowandy")
	functionPackageKeyword = []byte("/github.com/rizalgowandy/gdk/pkg/stack/")
	projectKeyword         = []byte("github.com/rizalgowandy/gdk")
)

// Trim removes unnecessary stack trace.
// Only take stack trace with current company keyword.
// Also, excludes lines with function package keyword.
func Trim(stack []byte) []byte {
	if stack == nil {
		return nil
	}

	// remove all log before panic keyword.
	idx := bytes.Index(stack, panicKeyword)
	if idx == -1 {
		idx = 0
	}
	stack = stack[idx:]

	// remove all log before company keyword.
	idx = bytes.Index(stack, companyKeyword)
	if idx == -1 {
		idx = 0
	}
	stack = stack[idx:]

	// remove all log before current function location keyword.
	idx = bytes.Index(stack, functionPackageKeyword)
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
// Input  => /home/rizalgowandy/go/src/github.com/rizalgowandy/gdk/pkg/stack.go 130
// Result => gdk/pkg/stack.go 130
func ToArr(stack []byte) []string {
	if stack == nil {
		return nil
	}

	msg := string(stack)
	arr := strings.Split(msg, "\n")

	var res []string
	for _, v := range arr {
		if v == "" {
			continue
		}

		tmp := []byte(v)
		idx := bytes.Index(tmp, projectKeyword)
		if idx == -1 {
			idx = 0
		}
		tmp = tmp[idx:]
		res = append(res, string(tmp))
	}

	return res
}
