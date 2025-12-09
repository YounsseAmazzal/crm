package database

import (
	"database/sql"
	"log"
	"main/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize database
func Init() *sql.DB {
	var err error

	// Do NOT use :=
	db, err = sql.Open("sqlite3", "./database/contacts.db")
	if err != nil {
		log.Fatal(err)
	}

	// Do NOT close the DB here!
	// defer db.Close()  ❌ REMOVE THIS

	createTableSQL := `CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT,
		phone TEXT,
		notes TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Selectfromdatabase() ([]models.Contact, error) {
	rows, err := db.Query("SELECT id, name, email, phone, notes FROM contacts ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact

	for rows.Next() {
		var c models.Contact
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Notes)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}
