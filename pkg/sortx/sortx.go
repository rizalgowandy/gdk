package sortx

import (
	"strings"
)

type (
	Key   string
	Order string
)

const (
	OrderAscending  Order = "ASC"
	OrderDescending Order = "DESC"
)

type Sort struct {
	Key      Key
	Order    Order
	Original string
}

// NewSorts create sorting based on
// Format:
//	sort=key1:asc,key2:desc,key3:asc
func NewSorts(qs string) []Sort {
	sorts := strings.Split(qs, ",")

	var res []Sort
	for _, v := range sorts {
		kv := strings.Split(v, ":")

		s := Sort{
			Key:      Key(strings.ToLower(kv[0])),
			Order:    OrderAscending,
			Original: strings.ToLower(v),
		}
		if len(kv) == 2 {
			switch strings.ToLower(kv[1]) {
			case "asc":
				s.Order = OrderAscending
			case "desc":
				s.Order = OrderDescending
			default:
				s.Order = OrderAscending
			}
		}
		res = append(res, s)
	}

	return res
}
