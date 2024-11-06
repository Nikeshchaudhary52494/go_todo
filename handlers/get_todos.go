package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Nikeshchaudhary52494/goTest/storage"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := storage.GetTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
