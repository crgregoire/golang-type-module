package types

import "net/http"

//
// Route is used to denote a route for all
// API routers
//
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
