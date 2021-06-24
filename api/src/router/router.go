package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Return a router with the routers
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
