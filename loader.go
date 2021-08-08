package main

import (
	_ "github.com/peractio/gdk/pkg/benchmark"
	_ "github.com/peractio/gdk/pkg/build"
	_ "github.com/peractio/gdk/pkg/converter"
	_ "github.com/peractio/gdk/pkg/cronx"
	_ "github.com/peractio/gdk/pkg/env"
	_ "github.com/peractio/gdk/pkg/errorx/v1"
	_ "github.com/peractio/gdk/pkg/errorx/v2"
	_ "github.com/peractio/gdk/pkg/httpx/echo"
	_ "github.com/peractio/gdk/pkg/httpx/echo/middleware"
	_ "github.com/peractio/gdk/pkg/httpx/mux"
	_ "github.com/peractio/gdk/pkg/httpx/mux/middleware"
	_ "github.com/peractio/gdk/pkg/logx"
	_ "github.com/peractio/gdk/pkg/password"
	_ "github.com/peractio/gdk/pkg/regex"
	_ "github.com/peractio/gdk/pkg/resync"
	_ "github.com/peractio/gdk/pkg/stack"
	_ "github.com/peractio/gdk/pkg/storage/cache"
	_ "github.com/peractio/gdk/pkg/storage/database"
	_ "github.com/peractio/gdk/pkg/tags"
	_ "github.com/peractio/gdk/pkg/ternary"
	_ "github.com/peractio/gdk/pkg/try"
	_ "github.com/peractio/gdk/pkg/validator"
)

// Ensure all package can be built correctly.
func main() {}
