package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	AuthMethod  int
}

const Admin = 0
const NoAuth = 0
const LoginNeeded = 0

type Routes []Route

var routes = Routes{
	Route{
		"StoreCreate",
		"GET",
		"/test",
		StoreCreate,
		Admin,
	},

}



const (
	NotAuthen = iota

)

func handleAuthenMethod(handler http.Handler, r Route) http.Handler {
	switch r.AuthMethod {
	case Admin:

	//case NoAuth:
	//
	//case LoginNeeded:
	}
	return handler
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		log.Infof("AuthMethod [%d]: %s\t%s\t%s", route.AuthMethod, route.Name, route.Method, route.Pattern)

		handler := handleAuthenMethod(route.HandlerFunc, route)
		router.
			//PathPrefix("/api/v1").
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
