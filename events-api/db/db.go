package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./events.db")

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	createTables()
}

func createTables() {

	createUsersTableQ := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)`

	_, err := DB.Exec(createUsersTableQ)

	if err != nil {
		panic("Failed to create users table")
	}

	createEventsTableQ := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date DATETIME NOT NULL,
			author_id INTEGER,
			FOREIGN KEY (author_id) REFERENCES users(id)
		)`

	_, err = DB.Exec(createEventsTableQ)

	if err != nil {
		panic("Failed to create events table")
	}

	createRegTableQ := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY (event_id) REFERENCES events(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`

	_, err = DB.Exec(createRegTableQ)

	if err != nil {
		panic("Failed to create registrations table")
	}
}
