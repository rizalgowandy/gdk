package regex

import (
	"regexp"
)

var (
	AlphanumericUnderScore      = regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
	AlphanumericUnderScoreSlash = regexp.MustCompile(`^[a-zA-Z0-9_/]*$`)
)
