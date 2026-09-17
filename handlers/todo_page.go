package handlers

import (
	"html/template"
	"net/http"
)

func TodoPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/todos" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Todos)

	if err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}
