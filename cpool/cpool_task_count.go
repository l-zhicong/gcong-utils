package cpool

import "sync/atomic"

type Count int64

var count Count

func DefaultCount() Count {
	atomic.AddInt64(count.int64(), 0)
	return 0
}

func (c *Count) int64() *int64 {
	return (*int64)(c)
}

func (c *Count) Add(count int) {
	atomic.AddInt64(c.int64(), int64(count))
}

func (c *Count) Cas(old, new int) (swapped bool) {
	return atomic.CompareAndSwapInt64(c.int64(), int64(old), int64(new))
}
