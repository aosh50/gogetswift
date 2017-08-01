package routing

import (
	"net/http"
	
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Test,
	},
	Route{
		"Drones",
		"GET",
		"/drones",
		Drones,
	},
	Route{
		"Packages",
		"GET",
		"/packages",
		Packages,
	},
	Route{
		"Update",
		"GET",
		"/update",
		Update,
	},
}
