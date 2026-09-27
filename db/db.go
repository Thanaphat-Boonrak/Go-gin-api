package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("sqlite3", "api")

	if err != nil {
		panic("Could not connect to database")
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	createTables()
}

func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY AUTOINCREMENT,email TEXT UNIQUE NOT NULL,password TEXT NOT NULL);`
	_, err := Db.Exec(createUserTable)
	if err != nil {
		panic("Could not create users table")
		return
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
)`

	_, err = Db.Exec(createEventTable)
	if err != nil {
		panic("Could not create events table")
		return
	}

	createRegisterTable := `CREATE TABLE IF NOT EXISTS registrations
	(id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE,
    UNIQUE (event_id, user_id)
	)`

	_, err = Db.Exec(createRegisterTable)
	if err != nil {
		panic("Could not create registrations table")
	}

}
