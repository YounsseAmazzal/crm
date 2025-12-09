package main

import (
	"log"
	"main/database"
	"main/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.Init()
	// Routes
	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/add", handlers.HandleAdd)
	//static file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
		log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}