package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rote
type Rote struct {
	URI                    string
	Method                 string
	Func                   func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

//Configuration of routes
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return r
}
