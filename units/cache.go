package units

import (
	"container/list"
	"fmt"
	"io"
	"sync"
	"time"
)

// go:generate go install golang.org/x/tools/cmd/stringer

//go:generate stringer -type MeasureUnit
type MeasureUnit int

func (m MeasureUnit) Compute(t time.Time) int64 {
	switch m {
	case NANOS:
		return t.UnixNano()
	case MILLS:
		return t.UnixMilli()
	case MICROS:
		return t.UnixMicro()
	case SECONDS:
		return t.Unix()
	default:
		panic(fmt.Errorf("unknown unit: %s", m))
	}
}
func (m MeasureUnit) ComputeTTL(t time.Duration) int64 {
	switch m {
	case NANOS:
		return t.Nanoseconds()
	case MILLS:
		return t.Milliseconds()
	case MICROS:
		return t.Microseconds()
	case SECONDS:
		return int64(t.Seconds())
	default:
		panic(fmt.Errorf("unknown unit: %s", m))
	}
}

const (
	NANOS MeasureUnit = iota
	MICROS
	MILLS
	SECONDS
)

//region LRU

type LRU[K comparable] struct {
	emptyKey K
	ll       *list.List
	m        map[K]*list.Element
}

func (l *LRU[K]) Size() int {
	return len(l.m)
}
func (l *LRU[K]) Add(k K) {
	e := l.ll.PushBack(k)
	l.m[k] = e
}
func (l *LRU[K]) Remove(k K) {
	e := l.m[k]
	l.ll.Remove(e)
	delete(l.m, k)
}
func (l *LRU[K]) Accessed(k K) {
	e := l.m[k]
	l.ll.MoveToBack(e)
}
func (l *LRU[K]) Oldest() K {
	e := l.ll.Front()
	if e == nil {
		return l.emptyKey
	}
	return e.Value.(K)
}
func (l *LRU[K]) Freshest() K {
	e := l.ll.Back()
	if e == nil {
		return l.emptyKey
	}
	return e.Value.(K)
}

func (l *LRU[K]) Clear() *LRU[K] {
	if l.ll == nil {
		l.ll = list.New()
	} else {
		l.ll.Init()
	}
	l.m = make(map[K]*list.Element)
	return l
}

// endregion

//region Entry

type Entry[K any, V any] struct {
	key    K
	data   V
	expire int64
}

func (e *Entry[K, V]) Key() K {
	return e.key
}
func (e *Entry[K, V]) Data() V {
	return e.data
}
func (e *Entry[K, V]) SetExpire(v int64) {
	e.expire = v
}
func (e *Entry[K, V]) Expire() int64 {
	return e.expire
}

// endregion
// region TTL

type TTL[K comparable, V any] struct {
	stop chan struct{}
	n    int
	tick int64
	time time.Duration
	kv   sync.Map
}

func (c *TTL[K, V]) addEntry(t *Entry[K, V]) {
	c.kv.Store(t.key, t)
	c.n += 1
}
func (c *TTL[K, V]) removeEntry(k K) {
	c.kv.Delete(k)
	c.n -= 1
}
func (c *TTL[K, V]) load(k K) (e *Entry[K, V], ok bool) {
	var v any
	v, ok = c.kv.Load(k)
	if ok {
		e = v.(*Entry[K, V])
	}
	return
}
func (c *TTL[K, V]) clear() {
	c.kv.Range(func(key, value any) bool {
		c.kv.Delete(key)
		return true
	})
}
func (c *TTL[K, V]) each(fn func(e *Entry[K, V]) bool) {
	c.kv.Range(func(key, value any) bool {
		return fn(value.(*Entry[K, V]))
	})
}
func (c *TTL[K, V]) HouseKeeping() bool {
	return c.stop != nil
}

func (c *TTL[K, V]) StopKeeping() {
	if c.stop == nil {
		return
	}
	c.stop <- struct{}{}

}
func (c *TTL[K, V]) Purify() {
	c.kv.Range(func(key, value any) bool {
		item := value.(*Entry[K, V])
		item.expire -= c.tick
		if item.expire <= 0 {
			c.kv.Delete(key)
			c.n -= 1
		}
		return true
	})
}
func (c *TTL[K, V]) Count() int {
	return c.n
}
func (c *TTL[K, V]) StartKeeping() {
	if c.stop != nil {
		return
	} else {
		c.stop = make(chan struct{})
	}
	go func() {
		ticker := time.NewTicker(c.time)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.Purify()
			case <-c.stop:
				close(c.stop)
				c.stop = nil
				return
			}
		}
	}()
}

//endregion
//region Cache

type Cache[K comparable, V any] interface {
	io.Closer
	EmptyValue() V                      //the empty value
	TimeToLive() time.Duration          // default time to live for entries
	MeasureUnit() MeasureUnit           //underlying measure time unit
	Put(k K, v V)                       //put value with default ttl
	PutTTL(k K, v V, ttl time.Duration) // put value with specified ttl
	Get(k K) (v V, ok bool)             //read a value by key
	Invalidate(k K)                     //remove entry by key
	InvalidateAll()                     //remove all entries
	Purify()                            //Purify entries
	HouseKeeping() bool                 // does housekeeping running
	StopKeeping()                       //close housekeeping
	StartKeeping()                      //start housekeeping
	All() map[K]V                       //dump all entries
	Count() int                         //current size
}
type cache[K comparable, V any] struct {
	empty V
	lru   *LRU[K]
	*TTL[K, V]
	limit      int
	timeToLive time.Duration
	ttl        int64
	unit       MeasureUnit
	onAccess   func(*Entry[K, V]) //!! change the entry when access
}

func (c *cache[K, V]) All() (m map[K]V) {
	m = make(map[K]V)
	c.each(func(e *Entry[K, V]) bool {
		m[e.key] = e.data
		return true
	})
	return
}
func (c *cache[K, V]) Invalidate(k K) {
	_, ok := c.load(k)
	if ok {
		c.removeEntry(k)
		if c.lru != nil {
			c.lru.Remove(k)
		}
	}
}
func (c *cache[K, V]) InvalidateAll() {
	c.clear()
	if c.lru != nil {
		c.lru.Clear()
	}
}
func (c *cache[K, V]) TimeToLive() time.Duration {
	return c.timeToLive
}
func (c *cache[K, V]) MeasureUnit() MeasureUnit {
	return c.unit
}
func (c *cache[K, V]) EmptyValue() V {
	return c.empty
}
func (c *cache[K, V]) Put(k K, v V) {
	itm := &Entry[K, V]{
		k,
		v,
		c.ttl,
	}
	c.addEntry(itm)
	if c.lru != nil {
		c.lru.Add(k)
		if c.limit > 0 && c.lru.Size() > c.limit {
			kx := c.lru.Oldest()
			c.lru.Remove(kx)
			c.removeEntry(kx)
		}
	}

}
func (c *cache[K, V]) PutTTL(k K, v V, ttl time.Duration) {
	itm := &Entry[K, V]{
		k,
		v,
		c.unit.Compute(time.Now().Add(ttl)),
	}
	c.addEntry(itm)
	if c.lru != nil {
		c.lru.Add(k)
		if c.limit > 0 && c.lru.Size() > c.limit {
			kx := c.lru.Oldest()
			c.lru.Remove(kx)
			c.removeEntry(kx)
		}
	}

}
func (c *cache[K, V]) Get(k K) (v V, ok bool) {
	e, ok := c.load(k)
	if !ok {
		return c.empty, false
	}
	if c.lru != nil {
		c.lru.Accessed(e.key)
	}
	if c.onAccess != nil {
		c.onAccess(e)
		if e.expire <= 0 {
			c.removeEntry(e.key)
			if c.lru != nil {
				c.lru.Remove(e.key)
			}
		}
	}
	return e.data, ok
}
func (c *cache[K, V]) Close() error {
	c.StopKeeping()
	return nil
}

//endregion

type Option = func(*conf)

// WithMaximize limit the maximum entries in Cache. also enable LRU inside Cache.
func WithMaximize(n uint32) Option {
	return func(c *conf) {
		c.max = int(n)
	}
}

// WithExpiredAfterAccess will change the entry expire time to specified duration after the access time
func WithExpiredAfterAccess(t time.Duration) Option {
	return func(c *conf) {
		c.onAccess = t
	}
}

type conf struct {
	max      int
	onAccess time.Duration
}

func NewCache[K comparable, V any](
	emptyKey K,
	emptyValue V,
	freq time.Duration,
	ttl time.Duration,
	unit MeasureUnit,
	opts ...Option,
) Cache[K, V] {
	c := new(conf)
	for _, opt := range opts {
		opt(c)
	}
	cc := new(cache[K, V])
	cc.empty = emptyValue

	cc.TTL = &TTL[K, V]{
		time: freq,
		tick: unit.ComputeTTL(freq),
		kv:   sync.Map{},
	}
	if c.max > 0 {
		cc.lru = (&LRU[K]{
			emptyKey: emptyKey,
		}).Clear()
		cc.limit = c.max
	}
	if c.onAccess != 0 {
		t := unit.ComputeTTL(c.onAccess)
		cc.onAccess = func(e *Entry[K, V]) {
			e.expire = t
		}
	}
	cc.timeToLive = ttl
	cc.ttl = unit.ComputeTTL(ttl)
	cc.unit = unit
	return cc
}
