package routes

import (
	"api/src/controller"
	"net/http"
)

var usersRoutes = []Rote{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   controller.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   controller.FindUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Func:                   controller.FindUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Func:                   controller.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Func:                   controller.DeleteUser,
		RequiresAuthentication: false,
	},
}
