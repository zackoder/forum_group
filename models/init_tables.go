package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitTables(db *sql.DB) {
	UsersTable(db)
	SessionsTable(db)
}

func runQuery(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func UsersTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`
	runQuery(db, query)
}

func SessionsTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS sessions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			passkey TEXT NOT NULL UNIQUE
		)
	`
	runQuery(db, query)
}
