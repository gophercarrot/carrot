package carrot

import (
	"sync"
)

//Counter struct for global increment
type Counter struct {
	val int
	mtx sync.Mutex
}

// Increment function can be called globally
func (counter *Counter) Increment() int {
	counter.mtx.Lock()
	counter.val++
	counter.mtx.Unlock()
	return counter.val
}
