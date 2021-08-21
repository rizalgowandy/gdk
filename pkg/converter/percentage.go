package converter

import (
	"fmt"
)

func Percentage(a, b interface{}) string {
	return fmt.Sprintf("%0.2f%%", Float64(a)*100/Float64(b))
}
