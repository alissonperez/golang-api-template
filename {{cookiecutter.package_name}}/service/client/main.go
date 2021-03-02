package client

import (
	"fmt"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/auth"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/repository"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/resources"
	"net/http"
)

type ClientService struct {
	repo repository.ClientRepository
	auth auth.Auth
}

func (s ClientService) GetClientFromRequest(r *http.Request) (resources.Client, error) {
	authData, err := s.auth.FromRequest(r)
	if err != nil {
		return resources.Client{}, err
	}

	client, ok := s.repo.GetById(authData.ClientId)
	if !ok {
		return resources.Client{}, fmt.Errorf("Client not found")
	}

	return client, nil
}

func NewService(repo repository.ClientRepository, auth auth.Auth) ClientService {
	return ClientService{repo: repo, auth: auth}
}
