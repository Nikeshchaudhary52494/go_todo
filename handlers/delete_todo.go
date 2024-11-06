package handlers

import (
	"net/http"
	"strconv"

	"github.com/Nikeshchaudhary52494/goTest/storage"
)

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if deleted := storage.DeleteTodoByID(id); !deleted {
		http.Error(w, "TODO not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
