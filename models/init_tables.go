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
		);
	`
	runQuery(db, query)
}

func PostsTable(db *sql.DB)               { // not complated
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER -- line not complated,
			title varchar(250) NOT NULL,
			content TEXT NOT NULL,
			image varchar(250),
			categories TEXT NOT NULL,
			date DATE NOT NULL
		);
	`
	runQuery(db, query)
}
func CategoriesTable(db *sql.DB)          {
	query := `
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name varchar(255) NOT NULL UNIQUE
		);
	`
	runQuery(db,query)
}
func CommentsTable(db *sql.DB)            { // not complated
	query := `
		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id -- line not complated,
			post_id -- line not complated,
			comment TEXT NOT NULL,
			date DATE NOT NULL
		);
	`
	runQuery(db,query)
}
func ReactionsTable(db *sql.DB)    { // not complated
	query := `
		CREATE TABLE IF NOT EXISTS reactions (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id -- line not complated,
			post_id -- line not complated,
			comment_id -- line not complated,
			type varchar(50) NOT NULL
		);
	`
	runQuery(db, query)
}
func PostsCategoriesTable(db *sql.DB) { // not complated
	query := `
		CREATE TABLE IF NOT EXISTS posts_categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			post_id -- line not complated,
			category_id -- line not complated,
		)
	`
	runQuery(db, query)
}
