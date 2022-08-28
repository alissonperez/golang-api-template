package teardown

import (
	"context"
	"go.uber.org/dig"
	"sync"
)

// Manages application teardown
type TearDown struct {
	context context.Context
	funcs   []*func()
}

func (t *TearDown) Register(downFunc *func()) {
	t.funcs = append(t.funcs, downFunc)
}

func (t *TearDown) Cancel() {
	waitGroup := sync.WaitGroup{}

	for _, downFunc := range t.funcs {
		waitGroup.Add(1)
		go func(downFunc *func()) {
			(*downFunc)()
			waitGroup.Done()
		}(downFunc)
	}

	waitGroup.Wait()
}

func newTearDown() *TearDown {
	td := TearDown{context: context.Background()}
	return &td
}

func Provide(container *dig.Container) {
	container.Provide(newTearDown)
}
