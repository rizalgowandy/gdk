package benchmark

import (
	"math/rand"
	"time"

	"github.com/peractio/gdk/pkg/converter"
	"github.com/pingcap/go-ycsb/pkg/generator"
)

type Hotspot struct {
	r *rand.Rand
	h *generator.Hotspot
}

// nolint:gosec
func NewHotspot(max int) Generator {
	return &Hotspot{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
		h: generator.NewHotspot(0, int64(max), 0.1, 0.9),
	}
}

func (g *Hotspot) Name() string {
	return "hostspot(0.1, 0.9)"
}

func (g *Hotspot) Next() string {
	return converter.String(g.h.Next(g.r))
}
