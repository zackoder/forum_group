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
	PostsTable(db)
	CategoriesTable(db)
	CommentsTable(db)
	ReactionsTable(db)
	PostsCategoriesTable(db)
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
			user_id INTEGER NOT NULL UNIQUE,
			passkey TEXT NOT NULL UNIQUE,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	runQuery(db, query)
}

func PostsTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title varchar(250) NOT NULL,
			content TEXT NOT NULL,
			image varchar(250),
			categories TEXT NOT NULL,
			date DATE NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	runQuery(db, query)
}

func CategoriesTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name varchar(255) NOT NULL UNIQUE
		);
	`
	runQuery(db, query)
}

func CommentsTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			comment TEXT NOT NULL,
			date DATE NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id)
		);
	`
	runQuery(db, query)
}

func ReactionsTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS reactions (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER,
			comment_id INTEGER,
			type varchar(50) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id),
			FOREIGN KEY (comment_id) REFERENCES comments(id),
			CHECK (
				(post_id IS NULL AND comment_id IS NOT NULL) OR
				(post_id IS NOT NULL AND comment_id IS NULL)
			)
		);
	`
	runQuery(db, query)
}

func PostsCategoriesTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS posts_categories (
			post_id INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			PRIMARY KEY (post_id,category_id),
			FOREIGN KEY (post_id) REFERENCES posts(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
		)
	`
	runQuery(db, query)
}
