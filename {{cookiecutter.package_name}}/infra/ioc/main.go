package ioc

import (
	"github.com/gorilla/mux"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/auth"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/config"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/infra/plog"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/infra/teardown"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/repository"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/service"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	c := dig.New()

	c.Provide(mux.NewRouter)

	teardown.Provide(c)
	config.Provide(c)
	auth.Provide(c)
	repository.Provide(c)
	service.Provide(c)
	plog.Provide(c)

	return c
}
