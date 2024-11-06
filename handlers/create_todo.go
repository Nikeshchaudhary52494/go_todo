package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Nikeshchaudhary52494/goTest/storage"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Task string `json:"task"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newTodo := storage.AddTodo(request.Task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}
