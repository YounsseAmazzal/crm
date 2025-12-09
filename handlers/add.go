package handlers

import (
	"database/sql"
	"net/http"
)
func DbConnection() *sql.DB {
	db, _ := sql.Open("sqlite3", "./database/contacts.db")
	return db
}
func HandleAdd(w http.ResponseWriter, r *http.Request) {
	db:=DbConnection()
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	notes := r.FormValue("notes")

	_, err := db.Exec("INSERT INTO contacts (name, email, phone, notes) VALUES (?, ?, ?, ?)",
		name, email, phone, notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}