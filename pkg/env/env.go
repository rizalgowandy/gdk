package env

import (
	"os"

	"github.com/peractio/gdk/pkg/syncx"
)

// List available environments.
const (
	Development = "development"
	Alpha       = "alpha"
	Beta        = "beta"
	Staging     = "staging"
	Production  = "production"
)

// Singleton pattern to prevent reading os more than once.
var (
	once       syncx.Once
	currentEnv string
)

// GetCurrent returns the current environment, if available.
// Otherwise returns environment as development.
func GetCurrent() string {
	once.Do(func() {
		env := os.Getenv("GDK_ENV")
		if env == "" {
			currentEnv = Development // set default as development
			return
		}
		currentEnv = env
	})
	return currentEnv
}

// IsDevelopment return true when current environment is development.
func IsDevelopment() bool {
	return GetCurrent() == Development
}

// IsAlpha return true when current environment is alpha.
func IsAlpha() bool {
	return GetCurrent() == Alpha
}

// IsBeta return true when current environment is beta.
func IsBeta() bool {
	return GetCurrent() == Beta
}

// IsStaging return true when current environment is staging.
func IsStaging() bool {
	return GetCurrent() == Staging
}

// IsProduction return true when current environment is production.
func IsProduction() bool {
	return GetCurrent() == Production
}
