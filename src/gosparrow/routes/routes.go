package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"gosparrow/app"
	"gosparrow/handlers"
)

// Route used to store a route item
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*app.Session, http.ResponseWriter, *http.Request) (
		int, interface{}, error)
}

// Routes is a list of Route
type Routes []Route

// NewRouter used to new a router object
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range gRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(formateHandler(route.HandlerFunc))
	}
	return router
}

// gRoutes is the global variable to store routes
var gRoutes = Routes{
	Route{
		"GetHeartBeat",
		"GET",
		"/heartbeat",
		handlers.Heartbeat,
	},
}

func formateHandler(f func(*app.Session, http.ResponseWriter, *http.Request) (
	int, interface{}, error)) http.HandlerFunc {

	return app.Handler(f).ServeHTTP
}
