package units

import (
	"context"
	"errors"
	"fmt"
	conf2 "github.com/ZenLiuCN/gofra/conf"
	"golang.org/x/exp/maps"
	"sync"
	"sync/atomic"
	"time"
)

type (

	// Scheduler compute the task to be scheduled.
	Scheduler[ID comparable, T comparable, V any] func(tick time.Time, v []*Task[ID, T, V])

	// Executor is a function that execute group of tasks
	Executor[ID comparable, T comparable, V any] func(tick time.Time, v []*Task[ID, T, V]) (failures []*Task[ID, T, V], err error)

	// Provider supplier of execute functions
	Provider[ID comparable, T comparable, V any] map[T]Executor[ID, T, V]

	// Task with one identity and one type and one value
	Task[ID comparable, T comparable, V any] struct {
		Stop  atomic.Bool
		Round atomic.Int64
		Id    ID
		Type  T
		Value V
	}
)

func (t *Task[ID, T, V]) GoString() string {
	return fmt.Sprintf("%v<%v>[%s](stop:%t,round:%d)", t.Id, t.Type, t.Value, t.Stop.Load(), t.Round.Load())
}

var (
	ErrBusy      = errors.New("registry busy")
	ErrNotExists = errors.New("not exists")
	ErrClosed    = errors.New("already closed")
)

// region Components
type entry[ID comparable, T comparable, V any] struct {
	Tasks []*Task[ID, T, V]
	When  uint64
	Stop  bool
}
type pool[ID comparable, T comparable, V any] struct {
	limit int
	s     sync.Pool
	m     sync.Pool
}

func (s *pool[ID, T, V]) init(limit, init int) *pool[ID, T, V] {
	s.limit = limit
	s.s.New = func() any {
		return make([]*Task[ID, T, V], 0, init)
	}
	s.m.New = func() any {
		return make(map[T][]*Task[ID, T, V])
	}
	return s
}
func (s *pool[ID, T, V]) GetMap() map[T][]*Task[ID, T, V] {
	return s.m.Get().(map[T][]*Task[ID, T, V])
}
func (s *pool[ID, T, V]) PutMap(v map[T][]*Task[ID, T, V]) {
	if len(v) > s.limit {
		return
	}
	maps.Clear(v)
	s.m.Put(v)
}

func (s *pool[ID, T, V]) GetList() []*Task[ID, T, V] {
	return s.s.Get().([]*Task[ID, T, V])
}
func (s *pool[ID, T, V]) PutList(v []*Task[ID, T, V]) {
	if len(v) > s.limit {
		return
	}
	v = v[:0]
	s.s.Put(v)
}

type action[ID comparable, T comparable, V any] struct {
	v []*Task[ID, T, V]
	t time.Time
}
type slots[ID comparable, T comparable, V any] struct {
	size     uint64
	tracef   func(string, ...any)
	interval time.Duration
	slot     atomic.Uint64
	event    chan entry[ID, T, V]
	tick     <-chan time.Time
	execute  chan action[ID, T, V]
	wheel    [][]*Task[ID, T, V]
	registry map[ID]*Task[ID, T, V]
	pool[ID, T, V]
}

// action listen and process execute channel
func (s *slots[ID, T, V]) action(ctx context.Context, submit func(tasks map[T][]*Task[ID, T, V], tk time.Time)) {
	for {
		select {
		case act, ok := <-s.execute:
			if !ok {
				s.tracef("execute chan closed")
				return
			}
			m := s.GetMap()
			for _, t := range act.v {
				//! ignore stopped
				if t.Stop.Load() {
					s.tracef("skip stopped action: %#v", t)
					continue
				}
				m[t.Type] = append(m[t.Type], t)
			}
			if len(m) == 0 {
				s.PutMap(m)
				s.tracef("skip empty actions")
				continue
			}
			s.tracef(" submit actions: %#+v", m)
			submit(m, act.t)
		case <-ctx.Done():
			s.tracef("context shutdown")
			//! not check
			close(s.event)
			s.event = nil
			return
		}
	}
}

// init initialize and start loop
func (s *slots[ID, T, V]) init(size, queueBuf, limit, init int, ticker <-chan time.Time) {
	s.size = uint64(size)
	s.event = make(chan entry[ID, T, V], queueBuf)
	s.wheel = make([][]*Task[ID, T, V], size)
	s.registry = make(map[ID]*Task[ID, T, V])
	s.pool.init(limit, init)
	s.execute = make(chan action[ID, T, V])
	s.tick = ticker
	go s.loop()
}
func (s *slots[ID, T, V]) reset(queueBuf int, ticker <-chan time.Time) {
	maps.Clear(s.registry)
	if s.event != nil {
		close(s.event)
	}
	s.wheel = make([][]*Task[ID, T, V], s.size)
	s.event = make(chan entry[ID, T, V], queueBuf)
	if s.execute != nil {
		close(s.execute)
	}
	s.execute = make(chan action[ID, T, V])
	s.tick = ticker
	go s.loop()
}
func (s *slots[ID, T, V]) register(v []*Task[ID, T, V], w uint64, awaitMax time.Duration) (err error) {
	if s.event == nil {
		return ErrClosed
	}
	c := time.NewTimer(awaitMax)
	defer c.Stop()
	for {
		select {
		case s.event <- entry[ID, T, V]{Tasks: v, When: w}:
			return nil
		case _ = <-c.C:
			s.tracef("register tasks timeout")
			return ErrBusy
		default:
		}
	}
}
func (s *slots[ID, T, V]) remove(v []*Task[ID, T, V], awaitMax time.Duration) (err error) {
	if s.event == nil {
		return ErrClosed
	}
	c := time.NewTimer(awaitMax)
	defer c.Stop()
	for {
		select {
		case s.event <- entry[ID, T, V]{Tasks: v, When: 0, Stop: true}:
			return nil
		case _ = <-c.C:
			s.tracef("remove tasks timeout")
			return ErrBusy
		default:
		}
	}
}
func (s *slots[ID, T, V]) loop() {
	for {
		select {
		case tick, ok := <-s.tick: //!! handle execute events
			if !ok {
				s.tracef("ticker closed")
				if s.execute != nil {
					close(s.execute)
					s.execute = nil
				}
				return
			}
			//! compute tasks
			p := s.slot.Add(1) % s.size
			if p == 0 {
				s.tracef("[0] reset slot")
				s.slot.Store(0)
			}
			v := s.wheel[p]  //! current tasks
			n := s.GetList() //! next round
			x := s.GetList() //! current execute
			s.tracef("[%d] scheduled tasks: %#+v", p, v)
			s.wheel[p] = s.GetList()
			for _, t := range v {
				//! ignore stopped task
				if t.Stop.Load() {
					delete(s.registry, t.Id)
					continue
				}
				//! dec circle
				r := t.Round.Add(-1)
				if r >= 0 {
					n = append(n, t)
				} else { //negative
					x = append(x, t)
				}
			}
			if len(n) > 0 {
				s.tracef("[%d] pushback tasks: %#+v", p, n)
				s.wheel[p] = append(s.wheel[p], n...)
			}
			s.PutList(n)
			s.PutList(v)
			if len(x) == 0 {
				s.tracef("[%d] ignore empty tasks", p)
				s.PutList(x)
				continue
			}
			s.tracef("[%d] unregister for executing tasks: %#+v", p, x)
			for _, t := range x {
				delete(s.registry, t.Id)
			}
			s.tracef("[%d] submit actions from tasks: %#+v", x)
			//! submit tasks
			s.execute <- action[ID, T, V]{x, tick}
		case e, ok := <-s.event: //!! handle register changes
			if !ok {
				if s.execute != nil {
					close(s.execute)
					s.execute = nil
				}
				return
			}
			if !e.Stop {
				//! register new tasks
				for _, task := range e.Tasks {
					if old, exist := s.registry[task.Id]; exist {
						s.tracef("conflict task: %#+v", old)
					}
					s.registry[task.Id] = task
				}
				s.tracef(" register tasks: %#+v", e.Tasks)
			}
			if e.Stop {
				//! set Stop mark only
				for _, task := range e.Tasks {
					task.Stop.Store(true)
					delete(s.registry, task.Id)
				}
				s.tracef("Stop tasks: %#+v", e.Tasks)
			} else {
				//! add to slot
				c := e.When / s.size //circles
				p := e.When % s.size //position
				if c > 0 {
					for _, task := range e.Tasks {
						task.Round.Store(int64(c))
					}
				}
				s.tracef(" schedule tasks: [%d:%d] %#+v", c, p, e.Tasks)
				s.wheel[p] = append(s.wheel[p], e.Tasks...)
			}
		}
	}
}

//endregion

type Ring[ID comparable, T comparable, V any] struct {
	noCopy
	retry    bool
	queueBuf int
	cc       func()
	trace    func(format string, args ...any)
	log      func(format string, args ...any)
	exec     Provider[ID, T, V]
	submit   func(func())
	slots[ID, T, V]
}
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
func (s *Ring[ID, T, V]) Tracef(format string, args ...any) {
	if s.trace != nil {
		s.trace(format, args...)
	}
}
func (s *Ring[ID, T, V]) Errorf(format string, args ...any) {
	if s.log != nil {
		s.log(format, args...)
	} else {
		conf2.Internal().Errorf(format, args...)
	}
}
func (s *Ring[ID, T, V]) start(ctx context.Context, size, queueBuf, limit, init int, ticker <-chan time.Time) {
	s.slots.init(size, queueBuf, limit, init, ticker)
	go s.slots.action(ctx, s.execute)
}

func (s *Ring[ID, T, V]) execute(tasks map[T][]*Task[ID, T, V], tk time.Time) {
	s.submit(func() {
		for i, t := range tasks {
			if x, ok := s.exec[i]; ok {
				s.submit(func() {
					f, err := x(tk, t)
					if err != nil {
						s.Errorf("execute error:%s at %s \n%#+v", err, tk, f)
					}
					if s.retry && len(f) > 0 {
						s.submit(func() {
							s.retrySchedule(f)
						})
					}
				})
			} else {
				s.submit(func() {
					s.retrySchedule(t)
				})
				s.Errorf("execute error:missing executor %v at %s \n%#+v", i, tk, t)
			}
		}
	})
}

var (
	RingRegisterMaxWait = time.Millisecond * 50
)

func (s *Ring[ID, T, V]) retrySchedule(f []*Task[ID, T, V]) {
	err := s.slots.register(f, s.slot.Load()+1, RingRegisterMaxWait)
	if err != nil {
		s.Errorf("retry schedule error:%s \n%#+v", err, f)
	}
}

// Register new task,if id is conflicted,the old ones may not be stopped. when is a relative tick location at current ring.
func (s *Ring[ID, T, V]) Register(id ID, kind T, value V, when uint64) (err error) {
	err = s.slots.register([]*Task[ID, T, V]{MakeTask(id, kind, value)}, when, RingRegisterMaxWait)
	return
}

// RegisterTasks like Register but do bach works.
func (s *Ring[ID, T, V]) RegisterTasks(when uint64, tasks ...*Task[ID, T, V]) (err error) {
	if len(tasks) == 0 {
		return
	}
	for _, task := range tasks { //! when user register a stopped task
		task.Stop.Store(false)
	}
	err = s.slots.register(tasks, when, RingRegisterMaxWait)
	return
}

// Remove tasks by id, if any of them haven't been registered returns ErrNotExists with extra error info.
func (s *Ring[ID, T, V]) Remove(id ...ID) (err error) {
	if len(id) == 0 {
		return nil
	}
	v := s.GetList()
	for _, i := range id {
		if t, ok := s.registry[i]; ok {
			v = append(v, t)
		} else {
			return errors.Join(ErrNotExists, fmt.Errorf("task of %v", i))
		}
	}
	err = s.slots.remove(v, RingRegisterMaxWait)
	return
}
func (s *Ring[ID, T, V]) RecycleTasks(v []*Task[ID, T, V]) {
	s.pool.PutList(v)
}

// Reset Stop current execution and clean all data. Then restart again.
func (s *Ring[ID, T, V]) Reset(ctx context.Context, ticker <-chan time.Time) {
	s.cc()
	ctx, s.cc = context.WithCancel(ctx)
	s.slots.reset(s.queueBuf, ticker)
	go s.slots.action(ctx, s.execute)
	return
}
func MakeTask[ID comparable, T comparable, V any](id ID, kind T, value V) *Task[ID, T, V] {
	return &Task[ID, T, V]{Id: id, Type: kind, Value: value}

}

/*
NewRing create new Ring. Ring is a ring buffered task executor. each tick will execute all task in current slot when the task is reached it's round.

ctx: required execute context;

retry: true to try in next tick;

trace: optional trace logging, this logger should skip 2 caller depth to reach real call point;

log: optional error logging, absent will use internal logger , this logger should skip 2 caller depth to reach real call point;

exec: map of every Type of Task to execute;

submit: optional goroutine pool, default just create new goroutines;

size: Ring slot size;

queueBuf: Ring event queue buffer size, this will limit the buffer task register events;

limit: Task list and Task Map limit size for [sync.Pool], only data lesser will be recycled;

init: Task list initialize size;

ticker: the Ticker to drive Ring;
*/
func NewRing[ID comparable, T comparable, V any](
	ctx context.Context,
	retry bool,
	trace func(format string, args ...any),
	log func(format string, args ...any),
	exec Provider[ID, T, V],
	submit func(func()),
	size, queueBuf, limit, init int,
	ticker <-chan time.Time) (x *Ring[ID, T, V]) {
	x = &Ring[ID, T, V]{}
	x.exec = exec
	if x.exec == nil || len(x.exec) == 0 {
		panic("Provider required!")
	}
	x.retry = retry
	x.trace = trace
	x.log = log
	x.queueBuf = queueBuf
	x.submit = submit
	if x.submit == nil {
		x.submit = func(f func()) {
			go f()
		}
	}
	x.slots.tracef = x.Tracef
	ctx, x.cc = context.WithCancel(ctx)
	x.start(ctx, size, queueBuf, limit, init, ticker)
	return
}
