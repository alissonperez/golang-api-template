package repository

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/resources"
)

type MemoryClientRepository struct {
	clients map[int]resources.Client
}

func (r MemoryClientRepository) Add(c resources.Client) {
	r.clients[c.Id] = c
}

func (r MemoryClientRepository) GetById(id int) (resources.Client, bool) {
	val, ok := r.clients[id]
	return val, ok
}

func CreateMemoryClientRepository() MemoryClientRepository {
	repo := MemoryClientRepository{}
	repo.clients = make(map[int]resources.Client)

	repo.Add(resources.Client{
		Id:       123,
		Name:     "Lojinha",
	})

	return repo
}
