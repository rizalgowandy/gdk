package fn

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const Company = "rizalgowandy"

// Name return the caller function name automatically.
// Format: package.struct.method
func Name(skips ...int) string {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}
	pc, _, _, _ := runtime.Caller(skip)
	return filepath.Base(runtime.FuncForPC(pc).Name())
}

// Line return the caller code line automatically.
// Format: /path/filename:line
func Line(skips ...int) string {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}
	_, file, line, _ := runtime.Caller(skip)
	file = file[strings.Index(file, Company)+len(Company):]
	return fmt.Sprintf("%s:%d", file, line)
}
