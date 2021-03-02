package v1

import (
	"net/http"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/net/v1/handler"
)

func WrapDependenciesHandler(handlerFunc func(w http.ResponseWriter, r *http.Request, deps *handler.HandlerDependencies), deps *handler.HandlerDependencies) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r, deps)
	}
}
