package converter

import (
	"strconv"
)

const (
	firstOrdinal  = 1
	secondOrdinal = 2
	thirdOrdinal  = 3

	firstOrdinalDivider  = 11
	secondOrdinalDivider = 12
	thirdOrdinalDivider  = 13
)

// Ordinal get ordinal 1st, 2nd, 3rd, etc.
func Ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case firstOrdinal:
		if x%100 != firstOrdinalDivider {
			suffix = "st"
		}
	case secondOrdinal:
		if x%100 != secondOrdinalDivider {
			suffix = "nd"
		}
	case thirdOrdinal:
		if x%100 != thirdOrdinalDivider {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
