package syncx

import (
	"sync/atomic"
	"time"
)

/**
 * Created by frankieci on 2022/2/17 4:43 pm
 */

// An AtomicDuration is an implementation of atomic duration.
type AtomicDuration int64

// NewAtomicDuration return a pointer type AtomicDuration
func NewAtomicDuration() *AtomicDuration {
	return new(AtomicDuration)
}

// ForAtomicDuration returns an AtomicDuration with given value.
func ForAtomicDuration(val time.Duration) *AtomicDuration {
	d := NewAtomicDuration()
	d.Set(val)
	return d
}

// Load loads the current duration.
func (d *AtomicDuration) Load() time.Duration {
	return time.Duration(atomic.LoadInt64((*int64)(d)))
}

// Set sets the value to val.
func (d *AtomicDuration) Set(val time.Duration) {
	atomic.StoreInt64((*int64)(d), int64(val))
}

// CompareAndSwap compares current value with old, if equals, set the value to val
func (d *AtomicDuration) CompareAndSwap(old, val time.Duration) bool {
	return atomic.CompareAndSwapInt64((*int64)(d), int64(old), int64(val))
}
