package handler

import (
	"go.uber.org/dig"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/infra/plog"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/service/client"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/service/url"
)

type HandlerDependencies struct {
	dig.In

	ClientService      client.ClientService
	UrlService         url.UrlService
	Logger             plog.Log
}
