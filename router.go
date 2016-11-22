package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	static := os.Getenv("STATIC")
	if static == "" {
		log.Fatal("You need to set STATIC environement variable, e.g. \"./client/build/\"")
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(static)))

	return router
}
