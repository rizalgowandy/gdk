package benchmark

import (
	"math/rand"
	"time"

	"github.com/peractio/gdk/pkg/converter"
	"github.com/pingcap/go-ycsb/pkg/generator"
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
	return converter.ToStr(g.h.Next(g.r))
}
