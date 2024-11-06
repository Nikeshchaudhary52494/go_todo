package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nikeshchaudhary52494/goTest/storage"
)

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	todo := storage.GetTodoByID(id)
	if todo == nil {
		http.Error(w, "TODO not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
