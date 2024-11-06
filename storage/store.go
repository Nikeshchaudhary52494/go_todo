package storage

import "github.com/Nikeshchaudhary52494/goTest/models"

var Todos = []models.Todo{}
var nextID = 1

func AddTodo(task string) models.Todo {
	todo := models.Todo{
		ID:        nextID,
		Task:      task,
		Completed: false,
	}
	Todos = append(Todos, todo)
	nextID++
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
			return &Todos[i]
		}
	}
	return nil
}

func DeleteTodoByID(id int) bool {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return true
		}
	}
	return false
}
