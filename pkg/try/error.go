package try

import (
	"errors"
)

var errMaxRetriesReached = errors.New("exceeded retry limit")
