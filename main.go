package main

import (
	"log"
	"net/http"

	"github.com/Nikeshchaudhary52494/goTest/auth"
	"github.com/Nikeshchaudhary52494/goTest/handlers"
	"github.com/Nikeshchaudhary52494/goTest/storage"
	"github.com/gorilla/mux"
)

func main() {
	if err := storage.LoadTodos(); err != nil {
		log.Fatalf("Failed to load todos: %v", err)
	}
	if err := storage.LoadUsers(); err != nil {
		log.Fatalf("Failed to load users: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

	todoRouter := r.PathPrefix("/todos").Subrouter()
	todoRouter.Use(auth.AuthMiddleware)

	todoRouter.HandleFunc("", handlers.GetTodosHandler).Methods("GET")
	todoRouter.HandleFunc("/create", handlers.CreateTodoHandler).Methods("POST")
	todoRouter.HandleFunc("/update", handlers.UpdateTodoHandler).Methods("PUT")
	todoRouter.HandleFunc("/delete", handlers.DeleteTodoHandler).Methods("DELETE")
	todoRouter.HandleFunc("/get", handlers.GetTodoHandler).Methods("GET")

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
