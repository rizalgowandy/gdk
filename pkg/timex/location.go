package timex

import (
	"sync"
	"time"
)

// JakartaLocation gives location information
var (
	once            sync.Once
	jakartaLocation *time.Location
)

// GetJakartaLocation return time location Asia/Jakarta WIB.
func GetJakartaLocation() *time.Location {
	once.Do(func() {
		var err error
		jakartaLocation, err = time.LoadLocation("Asia/Jakarta")
		if err != nil {
			// If LoadLocation from time Database failed, create the location with fixed value instead
			secondsEastOfUTC := int((7 * time.Hour).Seconds())

			// Use WIB as name so it can be parsed back into date.
			jakartaLocation = time.FixedZone("WIB", secondsEastOfUTC)
		}
	})

	return jakartaLocation
}
