package benchmark

import (
	"math/rand"
	"time"

	"github.com/pingcap/go-ycsb/pkg/generator"
	"github.com/rizalgowandy/gdk/pkg/converter"
)

type Uniform struct {
	r *rand.Rand
	h *generator.Uniform
}

// nolint:gosec
func NewUniform(max int) Generator {
	return &Uniform{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
		h: generator.NewUniform(0, int64(max)),
	}
}

func (g *Uniform) Name() string {
	return "uniform"
}

func (g *Uniform) Next() string {
	return converter.String(g.h.Next(g.r))
}
