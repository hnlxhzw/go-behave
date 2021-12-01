package decorator

import (
	"math/rand"
	"time"

	"github.com/woshihaomei/go-behave/core"
)

// RandomDelayer 随机延迟
func RandomDelayer(params core.Params, child core.Node) core.Node {
	base := core.NewDecorator("RandomDelayer", params, child)
	d := &randomDelayer{Decorator: base}

	msMin, err := params.GetInt("msMin")
	if err != nil {
		panic(err)
	}
	msMax, err := params.GetInt("msMax")
	if err != nil {
		panic(err)
	}

	d.delayMin = msMin
	d.delayMax = msMax
	return d
}

// delayer ...
type randomDelayer struct {
	*core.Decorator
	delayMin int
	delayMax int
	delay time.Duration // delay in milliseconds
	start time.Time
}

// Enter ...
func (d *randomDelayer) Enter(ctx *core.Context) {
	ms := rand.Intn(d.delayMax-d.delayMin) + d.delayMin
	d.delay = time.Duration(ms) * time.Millisecond
	d.start = time.Now()
}

// Tick ...
func (d *randomDelayer) Tick(ctx *core.Context) core.Status {
	if time.Since(d.start) > d.delay {
		return core.Update(d.Child, ctx)
	}
	return core.StatusRunning
}

// Leave ...
func (d *randomDelayer) Leave(ctx *core.Context) {}
