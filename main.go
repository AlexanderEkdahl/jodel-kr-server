package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatal("You need to set ADDR environement variable, e.g. \":8080\"")
	}

	log.Printf("Server started at %v", addr)

	r := NewRouter()
	log.Fatal(http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, handlers.CORS()(r))))
}
