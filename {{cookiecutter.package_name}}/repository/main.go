package repository

import (
	"go.uber.org/dig"
)

func Provide(container *dig.Container) {
	container.Provide(func() ClientRepository {
		return CreateMemoryClientRepository()
	})
}
