package pagination

//go:generate gomodifytags -all --quiet -w -file response.go -clear-tags
//go:generate gomodifytags -all --quiet --skip-unexported -w -file response.go -add-tags query,form,json,xml

type Response struct {
	Order         string  `query:"order" form:"order" json:"order" xml:"order"`
	StartingAfter *string `query:"starting_after" form:"starting_after" json:"starting_after" xml:"starting_after"`
	EndingBefore  *string `query:"ending_before" form:"ending_before" json:"ending_before" xml:"ending_before"`
	Total         int     `query:"total" form:"total" json:"total" xml:"total"`
	Yielded       int     `query:"yielded" form:"yielded" json:"yielded" xml:"yielded"`
	Limit         int     `query:"limit" form:"limit" json:"limit" xml:"limit"`
	PreviousURI   *string `query:"previous_uri" form:"previous_uri" json:"previous_uri" xml:"previous_uri"`
	NextURI       *string `query:"next_uri" form:"next_uri" json:"next_uri" xml:"next_uri"`
	// CursorRange returns cursors for starting after and ending before.
	// Format: [starting_after, ending_before].
	CursorRange []string `query:"cursor_range" form:"cursor_range" json:"cursor_range" xml:"cursor_range"`
}

// HasPrevPage returns true if prev page exists and can be traversed.
func (p *Response) HasPrevPage() bool {
	return p.PreviousURI != nil
}

// HasNextPage returns true if next page exists and can be traversed.
func (p *Response) HasNextPage() bool {
	return p.NextURI != nil
}

// PrevPageCursor returns cursor to be used as ending before value.
func (p *Response) PrevPageCursor() *string {
	if len(p.CursorRange) < 1 {
		return nil
	}
	return &p.CursorRange[0]
}

// NextPageCursor returns cursor to be used as starting after value.
func (p *Response) NextPageCursor() *string {
	if len(p.CursorRange) < 2 {
		return nil
	}
	return &p.CursorRange[1]
}

// PrevPageRequest returns pagination request for the prev page result.
func (p *Response) PrevPageRequest() Request {
	return Request{
		Order:         p.Order,
		Limit:         p.Limit,
		StartingAfter: nil,
		EndingBefore:  p.PrevPageCursor(),
	}
}

// NextPageRequest returns pagination request for the next page result.
func (p *Response) NextPageRequest() Request {
	return Request{
		Order:         p.Order,
		Limit:         p.Limit,
		StartingAfter: p.NextPageCursor(),
		EndingBefore:  nil,
	}
}
