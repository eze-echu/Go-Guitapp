package db

import (
	"database/sql"
	"github.com/eze-echu/guitapp/config"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(config config.DBConfig) {
	var err error
	DB, err = sql.Open(config.DbDriver, config.DbHost)
	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(config.DbMaxConns)
	DB.SetMaxIdleConns(config.DbMaxIdle)

	createTables()
}

func createTables() {
	createAccountTable := `
	CREATE TABLE IF NOT EXISTS accounts
	(
    account_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    current_value INTEGER NOT NULL DEFAULT 0,
    currency TEXT,
    rules INT
	)`
	_, err := DB.Exec(createAccountTable)
	if err != nil {
		panic("Failed to create users table")
	}
	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER,
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Failed to create events table")
	}
	createRegistrationTable := `CREATE TABLE IF NOT EXISTS registration (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(event_id) REFERENCES event(id)
)`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Failed to create registrations table")
	}

	log.Print("Table \"Events\" created")
}
