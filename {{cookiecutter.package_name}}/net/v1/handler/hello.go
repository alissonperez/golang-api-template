package handler

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/contrib/encoding"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/resources"
)

func Hello(w http.ResponseWriter, r *http.Request, deps *HandlerDependencies) {
	defer r.Body.Close()

	nameParam := mux.Vars(r)["myName"]

	client := r.Context().Value("client").(resources.Client)

	deps.Logger.Debugf("Saying hello to %s from client %s #%d", nameParam, client.Name, client.Id)

	result := struct {
		Message  string
		Name     string
		PathHere string
	}{
		Message:  fmt.Sprintf("Hello from client %s #%d...", client.Name, client.Id),
		Name:     nameParam,
		PathHere: deps.UrlService.Hello(nameParam),
	}

	b, err := json.Marshal(result)
	if err != nil {
		deps.Logger.Warnf("Error to marshall result: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	deps.Logger.Debugf("Response: %s", string(b))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoding.JsonEncode(w, result)
}
