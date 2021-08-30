package main

import (
	_ "github.com/rizalgowandy/gdk/pkg/benchmark"
	_ "github.com/rizalgowandy/gdk/pkg/build"
	_ "github.com/rizalgowandy/gdk/pkg/converter"
	_ "github.com/rizalgowandy/gdk/pkg/cronx"
	_ "github.com/rizalgowandy/gdk/pkg/cronx/interceptor"
	_ "github.com/rizalgowandy/gdk/pkg/env"
	_ "github.com/rizalgowandy/gdk/pkg/errorx/v1"
	_ "github.com/rizalgowandy/gdk/pkg/errorx/v2"
	_ "github.com/rizalgowandy/gdk/pkg/filepathx"
	_ "github.com/rizalgowandy/gdk/pkg/filex"
	_ "github.com/rizalgowandy/gdk/pkg/httpx"
	_ "github.com/rizalgowandy/gdk/pkg/httpx/echo"
	_ "github.com/rizalgowandy/gdk/pkg/httpx/echo/middleware"
	_ "github.com/rizalgowandy/gdk/pkg/httpx/mux"
	_ "github.com/rizalgowandy/gdk/pkg/httpx/mux/middleware"
	_ "github.com/rizalgowandy/gdk/pkg/jsonx"
	_ "github.com/rizalgowandy/gdk/pkg/logx"
	_ "github.com/rizalgowandy/gdk/pkg/netx"
	_ "github.com/rizalgowandy/gdk/pkg/password"
	_ "github.com/rizalgowandy/gdk/pkg/pprofx"
	_ "github.com/rizalgowandy/gdk/pkg/queue/nsqx"
	_ "github.com/rizalgowandy/gdk/pkg/queue/nsqx/interceptor"
	_ "github.com/rizalgowandy/gdk/pkg/regex"
	_ "github.com/rizalgowandy/gdk/pkg/stack"
	_ "github.com/rizalgowandy/gdk/pkg/storage/cache"
	_ "github.com/rizalgowandy/gdk/pkg/storage/database"
	_ "github.com/rizalgowandy/gdk/pkg/syncx"
	_ "github.com/rizalgowandy/gdk/pkg/tags"
	_ "github.com/rizalgowandy/gdk/pkg/ternary"
	_ "github.com/rizalgowandy/gdk/pkg/timex"
	_ "github.com/rizalgowandy/gdk/pkg/try"
	_ "github.com/rizalgowandy/gdk/pkg/validator"
)

// Ensure all package can be built correctly.
func main() {}
