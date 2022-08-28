package service

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/service/client"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/service/url"
	"go.uber.org/dig"
)

func Provide(container *dig.Container) {
	container.Provide(client.NewService)
	container.Provide(url.NewService)
}
