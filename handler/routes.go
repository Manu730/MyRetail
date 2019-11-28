package hndlr

import (
	"net/http"
)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var MyRetailRoutes = []Route{
	Route{
		"Get Products",
		"/api/v1/products/{id}",
		"GET",
		ServeProducts,
	},
	Route{
		"Update Products",
		"/api/v1/products/{id}",
		"PUT",
		UpdateProduct,
	},
	Route{
		"Products Update",
		"/api/v1/products/{id}",
		"POST",
		UpdateProduct,
	},
}
