package carrot

import (
	"sync"
)

//Counter struct for global increment
type Counter struct {
	val              int
	mtx              sync.Mutex
	success, failure int
}

// Increment function can be called globally
func (counter *Counter) Increment() int {
	counter.mtx.Lock()
	counter.val++
	counter.mtx.Unlock()
	return counter.val
}

func (counter *Counter) Success() int {
	counter.mtx.Lock()
	counter.success++
	counter.mtx.Unlock()
	return counter.success
}

func (counter *Counter) Failure() int {
	counter.mtx.Lock()
	counter.failure++
	counter.mtx.Unlock()
	return counter.failure
}
