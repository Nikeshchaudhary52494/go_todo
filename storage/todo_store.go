package storage

import (
	"encoding/json"
	"os"

	"github.com/Nikeshchaudhary52494/goTest/models"
)

var Todos []models.Todo
var nextID = 1

func LoadTodos() error {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			Todos = []models.Todo{}
			return nil
		}
		return err
	}

	if err := json.Unmarshal(data, &Todos); err != nil {
		return err
	}

	for _, todo := range Todos {
		if todo.ID >= nextID {
			nextID = todo.ID + 1
		}
	}
	return nil
}

func SaveTodos() error {
	data, err := json.MarshalIndent(Todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("todos.json", data, 0644)
}

func AddTodo(task string) models.Todo {
	todo := models.Todo{
		ID:        nextID,
		Task:      task,
		Completed: false,
	}
	Todos = append(Todos, todo)
	nextID++
	SaveTodos()
	return todo
}

func GetTodos() []models.Todo {
	return Todos
}

func GetTodoByID(id int) *models.Todo {
	for i, todo := range Todos {
		if todo.ID == id {
			return &Todos[i]
		}
	}
	return nil
}

func UpdateTodoByID(id int, task string, completed bool) *models.Todo {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos[i].Task = task
			Todos[i].Completed = completed
			SaveTodos()
			return &Todos[i]
		}
	}
	return nil
}

func DeleteTodoByID(id int) bool {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			SaveTodos()
			return true
		}
	}
	return false
}
