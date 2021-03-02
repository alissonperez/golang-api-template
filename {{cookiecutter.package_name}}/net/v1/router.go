package v1

import (
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"log"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/net/v1/handler"
)

func Setup(r *mux.Router, container *dig.Container) {
	err := container.Invoke(func(deps handler.HandlerDependencies) {
		// Get a transaction
		r.HandleFunc("/hello/{myName}", WrapDependenciesHandler(handler.Hello, &deps)).
			Methods("GET").
			Name("Hello")

		r.Use(WrapMiddleware(AuthMiddleware, &deps))
	})

	if err != nil {
		log.Fatalf("Error to invoke: %s", err)
	}
}
