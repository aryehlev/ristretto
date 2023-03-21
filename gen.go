package ristretto

import (
	"github.com/aryehlev/go-tinylfu"
	"math/rand"
	"time"

	"github.com/pingcap/go-ycsb/pkg/generator"
)

type NewGeneratorFunc = func(int) Generator

type Generator interface {
	Name() string
	Next() string
}

//------------------------------------------------------------------------------

type ScrambledZipfian struct {
	r *rand.Rand
	z *generator.ScrambledZipfian
}

func NewScrambledZipfian(max int) Generator {
	return &ScrambledZipfian{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
		z: generator.NewScrambledZipfian(0, int64(max), generator.ZipfianConstant),
	}
}

func (g *ScrambledZipfian) Name() string {
	return "zipfian"
}

func (g *ScrambledZipfian) Next() string {
	return go_tinylfu.stringFromInt64(g.z.Next(g.r))
}

//------------------------------------------------------------------------------

type Hotspot struct {
	r *rand.Rand
	h *generator.Hotspot
}

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
	return go_tinylfu.stringFromInt64(g.h.Next(g.r))
}

//------------------------------------------------------------------------------

type Uniform struct {
	r *rand.Rand
	h *generator.Uniform
}

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
	return go_tinylfu.stringFromInt64(g.h.Next(g.r))
}

type TestAryeh struct {
	i int64
	d int64
}

func NewTestAryeh(max int) Generator {
	return &TestAryeh{
		i: 0,
	}
}

func (g *TestAryeh) Name() string {
	return "aryeh"
}

func (g *TestAryeh) Next() string {
	if g.d%10 == 0 {
		g.i++
	}

	g.d++

	return go_tinylfu.stringFromInt64(g.i)
}
