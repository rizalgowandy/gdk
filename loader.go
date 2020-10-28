package main

import (
	_ "github.com/peractio/gdk/examples/myapp"
	_ "github.com/peractio/gdk/pkg/converter"
	_ "github.com/peractio/gdk/pkg/env"
	_ "github.com/peractio/gdk/pkg/errorx"
	_ "github.com/peractio/gdk/pkg/ternary"
	_ "github.com/peractio/gdk/pkg/try"
	_ "github.com/peractio/gdk/pkg/validator"
)

// Ensure all package can be built correctly.
func main() {}
