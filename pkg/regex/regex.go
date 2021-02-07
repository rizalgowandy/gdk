package regex

import (
	"regexp"
	"sync"
)

var storage sync.Map

// Register creates a new regex from pattern it has not been created at least once yet.
func Register(pattern string) (*regexp.Regexp, error) {
	if regex, ok := storage.Load(pattern); ok {
		return (regex).(*regexp.Regexp), nil
	}

	newMatcher, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	storage.Store(pattern, newMatcher)
	return newMatcher, nil
}

// MatchStr returns true if current string match the pattern, false otherwise.
func MatchStr(pattern, input string) (bool, error) {
	regex, ok := storage.Load(pattern)
	if !ok {
		newRegex, err := Register(pattern)
		if err != nil {
			return false, err
		}

		return newRegex.MatchString(input), nil
	}

	return (regex).(*regexp.Regexp).MatchString(input), nil
}
