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

// MatchString returns true if current string match the pattern, false otherwise.
func MatchString(pattern, input string) (bool, error) {
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

// ReplaceAllString returns a copy of input,
// replacing matches of the Regexp with the replacement string repl.
// Inside repl, $ signs are interpreted as in Expand,
// so for instance $1 represents the text of the first sub-match.
func ReplaceAllString(pattern, input, repl string) (string, error) {
	regex, ok := storage.Load(pattern)
	if !ok {
		newRegex, err := Register(pattern)
		if err != nil {
			return input, err
		}

		return newRegex.ReplaceAllString(input, repl), nil
	}

	return (regex).(*regexp.Regexp).ReplaceAllString(input, repl), nil
}
