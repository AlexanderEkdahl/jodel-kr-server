package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandlerGetMessages ...
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

// HandlerPostMessage ...
func HandlerPostMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	type params struct {
		Message string  `json:"message"`
		X       float32 `json:"x"`
		Y       float32 `json:"y"`
		UserID  string  `json:"user_id"`
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
		UserID:  p.UserID,
	}
	message, err = RepoCreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
}

// HandlerPostComment ...
func HandlerPostComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	type params struct {
		Content   string `json:"content"`
		MessageID int    `json:"message_id"`
		UserID    string `json:"user_id"`
	}
	var p params
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment := Comment{
		Content: p.Content,
		UserID:  p.UserID,
	}
	comment, err = RepoCreateComment(p.MessageID, comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

// HandlerGetMessagesWithUser ...
func HandlerGetMessagesWithUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	messages, err := RepoFindMessageWithUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
