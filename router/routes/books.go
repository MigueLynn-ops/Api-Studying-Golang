package routes

import (
	"net/http"

	"Api-Aula_1/controller"
)

var booksRoutes = []Routes{
	{
		URI:      "/books",
		Method:   http.MethodGet,
		Function: controller.HandleSearch,
	},
}
