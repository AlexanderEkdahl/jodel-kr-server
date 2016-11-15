package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("You need to set PORT environement variable")
	}

	log.Printf("Server started at http://localhost:%v", port)
	r := NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), handlers.CORS()(r)))
}
