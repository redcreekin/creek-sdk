package sdk

import (
	"context"
	stxlog "github.com/redcreekin/creek-sdk/log"
	"github.com/rockbears/log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

type GoRoutine struct {
	ctx     context.Context
	cancel  func()
	Name    string
	Func    func(ctx context.Context)
	Restart bool
	Active  bool
	mutex   sync.RWMutex
}

type GoRoutines struct {
	mutex  sync.RWMutex
	status []*GoRoutine
}

func (grs *GoRoutines) GoRoutine(name string) *GoRoutine {
	grs.mutex.Lock()
	defer grs.mutex.Unlock()
	for _, g := range grs.status {
		if g.Name == name {
			return g
		}
	}
	return nil
}

func (grs *GoRoutines) Stop(name string) {
	grs.mutex.Lock()
	defer grs.mutex.Unlock()
	for i, g := range grs.status {
		if g.Name == name {
			if g.cancel != nil {
				g.cancel()
			}
			grs.status = append(grs.status[:i], grs.status[i+1:]...)
			break
		}
	}
}

func (grs *GoRoutines) exec(g *GoRoutine) {
	hostname, _ := os.Hostname()

	go func(ctx context.Context) {
		ctx = context.WithValue(ctx, stxlog.Goroutine, g.Name)

		labels := pprof.Labels(
			"goroutine-name", g.Name,
			"goroutine-hostname", hostname,
			//"goroutine-id", fmt.Sprintf("%d", GoroutineID()),
		)
		goroutineCtx := pprof.WithLabels(ctx, labels)
		pprof.SetGoroutineLabels(goroutineCtx)

		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 1<<16)
				buf = buf[:runtime.Stack(buf, false)]
				ctx = context.WithValue(ctx, stxlog.Stacktrace, string(buf))
				log.Error(ctx, "[PANIC][%s] %s failed", hostname, g.Name)
			}
			g.mutex.Lock()
			g.Active = false
			g.mutex.Unlock()
		}()

		g.mutex.Lock()
		g.Active = true
		g.mutex.Unlock()
		g.Func(goroutineCtx)
	}(g.ctx)
}

func (grs *GoRoutines) Exec(ctx context.Context, name string, fn func(ctx context.Context)) {
	g := &GoRoutine{
		ctx:     ctx,
		Name:    name,
		Func:    fn,
		Active:  true,
		Restart: false,
	}
	grs.exec(g)
}

func NewGoRoutines(ctx context.Context) *GoRoutines {
	grs := &GoRoutines{}
	grs.Exec(ctx, "GoRoutines", func(ctx context.Context) {

	})
	return grs
}
