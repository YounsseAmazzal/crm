package handlers

import (
	"html/template"
	"main/database"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}

	contacts, err := database.Selectfromdatabase()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, contacts); err != nil {
		http.Error(w, "error rendering page", http.StatusInternalServerError)
	}
}
