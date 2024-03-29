package pagination

import (
	"github.com/rizalgowandy/gdk/pkg/converter"
)

//go:generate gomodifytags -all --quiet -w -file request.go -clear-tags
//go:generate gomodifytags -all --quiet --skip-unexported -w -file request.go -add-tags query,form,json,xml

// Request is a parameter to return list of data with pagination.
// Request is optional, most fields automatically filled by system.
// If you already have a response with pagination,
// you can generate pagination request directly to traverse next or prev page.
type Request struct {
	// Order of the resources in the response. desc (default), asc.
	// Order is optional.
	Order string `query:"order" form:"order" json:"order" xml:"order"`
	// Limit number of results per call. Accepted values: 0 - 100. Default 25
	// Limit is optional.
	Limit int `query:"limit" form:"limit" json:"limit" xml:"limit"`
	// StartingAfter is a cursor for use in pagination.
	// StartingAfter is a resource ID that defines your place in the list.
	// StartingAfter is optional.
	StartingAfter *string `query:"starting_after" form:"starting_after" json:"starting_after" xml:"starting_after"`
	// EndingBefore is cursor for use in pagination.
	// EndingBefore is a resource ID that defines your place in the list.
	// EndingBefore is optional.
	EndingBefore *string `query:"ending_before" form:"ending_before" json:"ending_before" xml:"ending_before"`
}

func (p Request) QueryParams() map[string]string {
	res := map[string]string{}
	if p.Order != "" {
		res["order"] = p.Order
	}
	if p.Limit > 0 {
		res["limit"] = converter.String(p.Limit)
	}
	if p.StartingAfter != nil {
		res["starting_after"] = *p.StartingAfter
	}
	if p.EndingBefore != nil {
		res["ending_before"] = *p.EndingBefore
	}
	return res
}
