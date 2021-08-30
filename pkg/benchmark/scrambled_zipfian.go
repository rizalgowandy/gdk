package benchmark

import (
	"math/rand"
	"time"

	"github.com/pingcap/go-ycsb/pkg/generator"
	"github.com/rizalgowandy/gdk/pkg/converter"
)

type ScrambledZipfian struct {
	r *rand.Rand
	z *generator.ScrambledZipfian
}

// nolint:gosec
func NewScrambledZipfian(max int) Generator {
	return &ScrambledZipfian{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
		z: generator.NewScrambledZipfian(0, int64(max), generator.ZipfianConstant),
	}
}

func (g *ScrambledZipfian) Name() string {
	return "scrambled_zipfian"
}

func (g *ScrambledZipfian) Next() string {
	return converter.String(g.z.Next(g.r))
}
