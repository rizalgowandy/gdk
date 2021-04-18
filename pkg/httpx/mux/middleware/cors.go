package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/peractio/gdk/pkg/env"
	"github.com/peractio/gdk/pkg/httpx/mux"
)

const maxDomainLength = 253

type (
	// CORSConfig defines the config for CORS middleware.
	CORSConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper Skipper

		// AllowOrigin defines a list of origins that may access the resource.
		// Optional. Default value []string{"*"}.
		AllowOrigins []string `yaml:"allow_origins"`

		// AllowMethods defines a list methods allowed when accessing the resource.
		// This is used in response to a preflight request.
		// Optional. Default value DefaultCORSConfig.AllowMethods.
		AllowMethods []string `yaml:"allow_methods"`

		// AllowHeaders defines a list of request headers that can be used when
		// making the actual request. This is in response to a preflight request.
		// Optional. Default value []string{}.
		AllowHeaders []string `yaml:"allow_headers"`

		// AllowCredentials indicates whether or not the response to the request
		// can be exposed when the credentials flag is true. When used as part of
		// a response to a preflight request, this indicates whether or not the
		// actual request can be made using credentials.
		// Optional. Default value false.
		AllowCredentials bool `yaml:"allow_credentials"`

		// ExposeHeaders defines a whitelist headers that clients are allowed to
		// access.
		// Optional. Default value []string{}.
		ExposeHeaders []string `yaml:"expose_headers"`

		// MaxAge indicates how long (in seconds) the results of a preflight request
		// can be cached.
		// Optional. Default value 0.
		MaxAge int `yaml:"max_age"`
	}
)

var (
	// DefaultCORSConfig is the default CORS middleware config.
	DefaultCORSConfig = CORSConfig{
		Skipper:      DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}
)

// CORS returns a Cross-Origin Resource Sharing (CORS) middleware.
// See: https://developer.mozilla.org/en/docs/Web/HTTP/Access_control_CORS
func CORS(next http.Handler) http.Handler {
	// for [development] then allow all origin.
	if env.IsDevelopment() {
		return CORSWithConfig(next, &DefaultCORSConfig)
	}

	// for [staging,beta,uat,production] enforce domain.
	return CORSWithConfig(next, &CORSConfig{
		Skipper: func(res http.ResponseWriter, req *http.Request) bool {
			return false
		},
		AllowOrigins: []string{
			"*.peract.io",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowHeaders: []string{
			mux.HeaderContentType,
		},
		AllowCredentials: true,
		ExposeHeaders:    nil,
		MaxAge:           3600,
	})
}

// CORSWithConfig returns a Cross-Origin Resource Sharing (CORS) middleware.
// nolint gocyclo
func CORSWithConfig(next http.Handler, config *CORSConfig) http.Handler {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultCORSConfig.Skipper
	}
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = DefaultCORSConfig.AllowOrigins
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = DefaultCORSConfig.AllowMethods
	}

	allowMethods := strings.Join(config.AllowMethods, ",")
	allowHeaders := strings.Join(config.AllowHeaders, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")
	maxAge := strconv.Itoa(config.MaxAge)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if config.Skipper(res, req) {
			next.ServeHTTP(res, req)
			return
		}

		origin := req.Header.Get(mux.HeaderOrigin)
		allowOrigin := ""

		// Check allowed origins
		for _, o := range config.AllowOrigins {
			if o == "*" && config.AllowCredentials {
				allowOrigin = origin
				break
			}
			if o == "*" || o == origin {
				allowOrigin = o
				break
			}
			if matchSubdomain(origin, o) {
				allowOrigin = origin
				break
			}
		}

		// Simple request
		if req.Method != http.MethodOptions {
			res.Header().Add(mux.HeaderVary, mux.HeaderOrigin)
			res.Header().Set(mux.HeaderAccessControlAllowOrigin, allowOrigin)
			if config.AllowCredentials {
				res.Header().Set(mux.HeaderAccessControlAllowCredentials, "true")
			}
			if exposeHeaders != "" {
				res.Header().Set(mux.HeaderAccessControlExposeHeaders, exposeHeaders)
			}
			next.ServeHTTP(res, req)
			return
		}

		// Preflight request
		res.Header().Add(mux.HeaderVary, mux.HeaderOrigin)
		res.Header().Add(mux.HeaderVary, mux.HeaderAccessControlRequestMethod)
		res.Header().Add(mux.HeaderVary, mux.HeaderAccessControlRequestHeaders)
		res.Header().Set(mux.HeaderAccessControlAllowOrigin, allowOrigin)
		res.Header().Set(mux.HeaderAccessControlAllowMethods, allowMethods)
		if config.AllowCredentials {
			res.Header().Set(mux.HeaderAccessControlAllowCredentials, "true")
		}
		if allowHeaders != "" {
			res.Header().Set(mux.HeaderAccessControlAllowHeaders, allowHeaders)
		} else {
			h := req.Header.Get(mux.HeaderAccessControlRequestHeaders)
			if h != "" {
				res.Header().Set(mux.HeaderAccessControlAllowHeaders, h)
			}
		}
		if config.MaxAge > 0 {
			res.Header().Set(mux.HeaderAccessControlMaxAge, maxAge)
		}

		res.WriteHeader(http.StatusNoContent)
	})
}

func matchScheme(domain, pattern string) bool {
	didx := strings.Index(domain, ":")
	pidx := strings.Index(pattern, ":")
	return didx != -1 && pidx != -1 && domain[:didx] == pattern[:pidx]
}

// matchSubdomain compares authority with wildcard
func matchSubdomain(domain, pattern string) bool {
	if !matchScheme(domain, pattern) {
		return false
	}

	didx := strings.Index(domain, "://")
	pidx := strings.Index(pattern, "://")
	if didx == -1 || pidx == -1 {
		return false
	}

	domAuth := domain[didx+3:]
	// Avoid long loop by invalid long domain.
	if len(domAuth) > maxDomainLength {
		return false
	}
	patAuth := pattern[pidx+3:]

	domComp := strings.Split(domAuth, ".")
	patComp := strings.Split(patAuth, ".")
	for i := len(domComp)/2 - 1; i >= 0; i-- {
		opp := len(domComp) - 1 - i
		domComp[i], domComp[opp] = domComp[opp], domComp[i]
	}
	for i := len(patComp)/2 - 1; i >= 0; i-- {
		opp := len(patComp) - 1 - i
		patComp[i], patComp[opp] = patComp[opp], patComp[i]
	}

	for i, v := range domComp {
		if len(patComp) <= i {
			return false
		}
		p := patComp[i]
		if p == "*" {
			return true
		}
		if p != v {
			return false
		}
	}
	return false
}
