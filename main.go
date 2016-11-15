package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("You need to set PORT environement variable")
	}

	log.Printf("Server started at http://localhost:%v", port)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
