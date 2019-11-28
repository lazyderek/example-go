package waitgroup

import "sync"

type WaitGroup struct {
	sync.WaitGroup
}

func New() *WaitGroup {
	return &WaitGroup{}
}

func (w *WaitGroup) Wrap(fn func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		fn()
	}()
}
