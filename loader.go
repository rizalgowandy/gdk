package main

import (
	_ "github.com/peractio/gdk/examples/myapp"
	_ "github.com/peractio/gdk/pkg/converter"
	_ "github.com/peractio/gdk/pkg/regex"
	_ "github.com/peractio/gdk/pkg/validator"
)

// Ensure all package can be built correctly.
func main() {}
