package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Nikeshchaudhary52494/goTest/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", handlers.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todos/create", handlers.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos/update", handlers.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/delete", handlers.DeleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/todos/get", handlers.GetTodoHandler).Methods("GET")

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
