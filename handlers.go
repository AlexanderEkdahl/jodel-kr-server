package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func HandlerGetMessages(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages, err := RepoFindMessage(float32(x), float32(y))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := Message{
		Message: p.Message,
		X:       p.X,
		Y:       p.Y,
	}
	message, err = RepoCreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
}
