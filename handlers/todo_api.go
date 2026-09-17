package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"to-do-list/To-Do-List/models"
)

var Todos []models.Todo
var NextID = 1

func GetTodosAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todos)
}

func CreateTodoAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	todo.Title = strings.TrimSpace(todo.Title)

	if todo.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	todo.ID = NextID
	todo.Completed = false
	NextID++

	Todos = append(Todos, todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
