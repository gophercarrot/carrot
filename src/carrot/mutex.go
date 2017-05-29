package carrot

import (
	"sync"
)

type Counter struct {
	val int
	mtx sync.Mutex
}

func (counter *Counter) Increment() int {
	counter.mtx.Lock()
	counter.val++
	counter.mtx.Unlock()
	return counter.val
}
