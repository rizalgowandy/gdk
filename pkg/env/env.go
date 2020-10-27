package env

import (
	"os"
)

// List available environments.
const (
	Development = "development"
	Alpha       = "alpha"
	Beta        = "beta"
	Staging     = "staging"
	Production  = "production"
)

// Get returns the current environment, if available.
// Otherwise returns environment as development.
func Get() string {
	env := os.Getenv("GDK_ENV")
	if env != "" {
		return env
	}

	return Development
}

// IsDevelopment return true when current environment is development.
func IsDevelopment() bool {
	return Get() == Development
}

// IsAlpha return true when current environment is alpha.
func IsAlpha() bool {
	return Get() == Alpha
}

// IsBeta return true when current environment is beta.
func IsBeta() bool {
	return Get() == Beta
}

// IsStaging return true when current environment is staging.
func IsStaging() bool {
	return Get() == Staging
}

// IsProduction return true when current environment is production.
func IsProduction() bool {
	return Get() == Production
}
