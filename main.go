package main

import (
	"log"
	"net/http"
	"to-do-list/handlers"
)

func main() {
	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetTodosAPI(w, r)
			return
		}
		if r.Method == http.MethodPost {
			handlers.CreateTodoAPI(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
	http.HandleFunc("/api/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handlers.UpdateTodoAPI(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			handlers.DeleteTodoAPI(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/todos", handlers.TodoPage)

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	log.Println("Server running at http://localhost:8080/todos")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
