package syncx

import "sync"

/**
 * Created by frankieci on 2022/3/3 2:03 pm
 */

// Barrier is used to facility the barrier on a resource
type Barrier struct {
	lock sync.Mutex
}

// Guard guards the given fn on the resource.
func (b *Barrier) Guard(fn func()) {
	Guard(&b.lock, fn)
}

// Guard guards the given fn with lock.
func Guard(lock sync.Locker, fn func()) {
	lock.Lock()
	defer lock.Unlock()
	fn()
}
