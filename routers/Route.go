package routers

import (
	"net/http"
)

/*Routes - routes */
type Routes []Route

/*Route - route */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
