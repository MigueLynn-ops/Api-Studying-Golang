package router

import (
	"Api-Aula_1/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	routes.Register(r)
	return r
}
