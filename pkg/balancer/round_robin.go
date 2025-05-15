package balancer

import (
	"sync"

	"github.com/rizalgowandy/gdk/pkg/errorx/v1"
)

func NewRoundRobin(items []any) (*RoundRobin, error) {
	if len(items) == 0 {
		return nil, errorx.E("no items passed")
	}

	return &RoundRobin{
		items: items,
	}, nil
}

type RoundRobin struct {
	mux   sync.Mutex
	next  int
	items []any
}

func (b *RoundRobin) Next() any {
	b.mux.Lock()
	r := b.items[b.next]
	b.next = (b.next + 1) % len(b.items)
	b.mux.Unlock()
	return r
}
