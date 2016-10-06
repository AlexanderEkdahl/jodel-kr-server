package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func HandlerGetMessages(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 32)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 32)
	if err != nil {
		log.Fatal(err)
	}

	messages := RepoFindMessage(float32(x), float32(y))

	json.NewEncoder(w).Encode(messages)
}

func HandlerPostMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	x, err := strconv.ParseFloat(r.Form.Get("x"), 32)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseFloat(r.Form.Get("y"), 32)
	if err != nil {
		log.Fatal(err)
	}

	message := Message{
		Message: r.Form.Get("message"),
		X:       float32(x),
		Y:       float32(y),
	}
	message = RepoCreateMessage(message)
	json.NewEncoder(w).Encode(message)
}
