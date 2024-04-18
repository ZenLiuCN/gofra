package utils

import (
	"bytes"
	"sync"
)

type Pool[T any] struct {
	p     sync.Pool
	reset func(T) T
}

func NewPool[T any](ctor func() T, reset func(T) T) *Pool[T] {
	return &Pool[T]{p: sync.Pool{New: func() any { return ctor() }}, reset: reset}
}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}
func (p *Pool[T]) Put(t T) {
	p.p.Put(p.reset(t))
}
func NewByteBufferPool() *Pool[*bytes.Buffer] {
	return NewPool(func() *bytes.Buffer {
		return new(bytes.Buffer)
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	})
}
