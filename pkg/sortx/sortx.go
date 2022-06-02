package sortx

import (
	"fmt"
	"strings"
)

const (
	OrderAscending  Order = "ASC"
	OrderDescending Order = "DESC"
)

type Key string

func (k Key) String() string {
	return string(k)
}

type Order string

func (o Order) String() string {
	return string(o)
}

type Sort struct {
	Key      Key
	Order    Order
	Original string
}

// NewSorts create sorting based on
// Format:
//	sort=key1:asc,key2:desc,key3:asc
func NewSorts(qs string) Sorts {
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

type Sorts []Sort

func (s Sorts) OrderBy(reverse ...bool) string {
	var res string
	for _, v := range s {
		cur := v.Order
		if len(reverse) > 0 && reverse[0] {
			switch v.Order {
			case OrderAscending:
				cur = OrderDescending
			case OrderDescending:
				cur = OrderAscending
			}
		}
		if res != "" {
			res += ", "
		}
		res += fmt.Sprintf("%s %s", v.Key, cur)
	}
	return res
}

func (s Sorts) Map() map[string]string {
	res := make(map[string]string)
	for _, v := range s {
		res[v.Key.String()] = v.Order.String()
	}
	return res
}
