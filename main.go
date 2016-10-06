package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started at http://localhost:8080")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
