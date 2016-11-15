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
	decoder := json.NewDecoder(r.Body)

	type params struct {
		Message string  `json:"message"`
		X       float32 `json:"x"`
		Y       float32 `json:"y"`
	}
	var p params
	err := decoder.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	message := Message{
		Message: p.Message,
		X:       p.X,
		Y:       p.Y,
	}
	message = RepoCreateMessage(message)
	json.NewEncoder(w).Encode(message)
}
