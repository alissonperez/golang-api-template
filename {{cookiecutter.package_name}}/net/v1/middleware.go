package v1

import (
	"context"
	"fmt"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/net/v1/handler"
	"net/http"
)

func AuthMiddleware(next http.Handler, deps *handler.HandlerDependencies) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := deps.ClientService.GetClientFromRequest(r)
		if err != nil {
			deps.Logger.Warnf("Error to get client from request: %s", err)
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Client not found/Invalid Credentials")
			return
		}

		deps.Logger.Debugf("Client %s #%d found by request", client.Name, client.Id)

		ctx := context.WithValue(r.Context(), "client", client)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WrapMiddleware(middleware func(next http.Handler, deps *handler.HandlerDependencies) http.Handler, deps *handler.HandlerDependencies) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return middleware(next, deps)
	}
}
