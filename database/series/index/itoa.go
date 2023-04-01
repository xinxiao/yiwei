package index

import (
	"sync"
)

type itoa struct {
	curr int64
	mut  sync.Mutex
}

func (i *itoa) Generate() int64 {
	i.mut.Lock()
	defer i.mut.Unlock()

	curr := i.curr
	i.curr += 1
	return curr
}
