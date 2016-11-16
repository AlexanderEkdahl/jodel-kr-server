package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	static := os.Getenv("STATIC")
	if static == "" {
		log.Fatal("You need to set STATIC environement variable, e.g. \"./client/build/\"")
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(static)))

	return router
}
