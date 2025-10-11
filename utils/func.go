package utils

import (
	"database/sql"
	"log"
)

// Must Takes care of boilerplate error code.
func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// InitSchema Create DB if not existing already.
func InitSchema(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS items (
  		name TEXT PRIMARY KEY,
  		unit_cost REAL NOT NULL DEFAULT 0,
  		unit_store REAL NOT NULL DEFAULT 0,
  		unit TEXT NOT NULL DEFAULT ''
);
	CREATE TABLE IF NOT EXISTS bakes (
  		id INTEGER PRIMARY KEY AUTOINCREMENT,
  		when_at TEXT NOT NULL,
  		item_name TEXT NOT NULL,
  		qty INTEGER NOT NULL,
  		user TEXT NOT NULL DEFAULT '',
  		FOREIGN KEY(item_name) REFERENCES items(name)
);
	`)
	return err // return any errors
}
