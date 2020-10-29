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

// GetCurrent returns the current environment, if available.
// Otherwise returns environment as development.
func GetCurrent() string {
	env := os.Getenv("GDK_ENV")
	if env != "" {
		return env
	}

	return Development
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
