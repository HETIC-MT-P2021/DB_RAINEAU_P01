package router

import (
	"github.com/gorilla/mux"
	"gobdd/controllers"
	"net/http"
)

// Route struct defining all of this project routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice of Route
type Routes []Route

// newRouter registers public routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:        "Home",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.RenderHome,
	},
	Route{
		Name:        "Get a customer by id",
		Method:      "GET",
		Pattern:     "/customer/{id}",
		HandlerFunc: controllers.RenderCustomer,
	},
	Route{
		Name:        "Get a command by id",
		Method:      "GET",
		Pattern:     "/order/{id}",
		HandlerFunc: controllers.RenderOrder,
	},
	Route{
		Name: 		"Get all employes",
		Method:      "GET",
		Pattern:     "/employees",
		HandlerFunc: controllers.RenderEmployees,
	},
}
