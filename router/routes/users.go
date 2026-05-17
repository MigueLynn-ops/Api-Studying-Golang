package routes

import (
	"Api-Aula_1/controller"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controller.CreateUser,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controller.FetchUser,
	},
	{
		URI:      "/users{userID}",
		Method:   http.MethodGet,
		Function: controller.UpdateUser,
	},
	{
		URI:      "/users{userID}",
		Method:   http.MethodDelete,
		Function: controller.DeleteUser,
	},
}
