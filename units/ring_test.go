package units

import (
	"context"
	"flag"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/golang/glog"
	"os"
	"testing"
	"time"
)

type SimpleRing struct {
	*Ring[string, int, string]
}

func TestRing(t *testing.T) {
	os.Args = append(os.Args, "-alsologtostderr")
	flag.Parse()
	now := time.Now()
	exe := map[string]time.Duration{}
	ctx, cc := context.WithTimeout(context.Background(), time.Second*150)
	defer cc()
	var ring SimpleRing
	ring = SimpleRing{
		NewRing(ctx, true, func(format string, args ...any) {
			glog.InfoDepthf(2, format, args...)
		}, func(format string, args ...any) {
			glog.ErrorDepthf(2, format, args...)
		}, map[int]Executor[string, int, string]{
			0: func(tick time.Time, v []*Task[string, int, string]) (failures []*Task[string, int, string], err error) {
				t.Logf("%s on tasks: %#+v", tick, v)
				for _, t2 := range v {
					exe[t2.Id] = tick.Sub(now)
				}
				ring.RecycleTasks(v)
				return nil, nil
			},
		}, nil, 60, 5, 60, 5, time.Tick(time.Second)),
	}
	for i := 0; i < 60; i++ {
		val := fmt.Sprintf("%d", i)
		v := i
		fn.Panic(ring.Register(val, 0, val, uint64(v+1)))
	}
	for i := 60; i < 120; i++ {
		val := fmt.Sprintf("%d", i)
		v := i
		fn.Panic(ring.Register(val, 0, val, uint64(v)))
	}
	<-ctx.Done()
	t.Logf("%#+v", ring.registry)
	t.Logf("%s", ctx.Err())
	t.Logf("total %d", len(exe))
	for s, duration := range exe {
		t.Logf("%s at %s", s, duration)
	}
}
