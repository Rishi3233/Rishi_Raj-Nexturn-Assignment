package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	// Open a SQLite database file
	DB, err = sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the 'products' table if it doesn't already exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price FLOAT NOT NULL,
			stock INTEGER NOT NULL,
			category_id INTEGER
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
