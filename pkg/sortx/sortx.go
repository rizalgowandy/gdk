package sortx

import "strings"

type (
	Key   string
	Order int64
)

const (
	SortOrderAscending Order = iota + 1
	SortOrderDescending
)

type Sort struct {
	Key   Key
	Order Order
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
			Key:   Key(kv[0]),
			Order: SortOrderAscending,
		}
		if len(kv) == 2 {
			switch kv[1] {
			case "asc":
				s.Order = SortOrderAscending
			case "desc":
				s.Order = SortOrderDescending
			default:
				s.Order = SortOrderAscending
			}
		}
		res = append(res, s)
	}

	return res
}
