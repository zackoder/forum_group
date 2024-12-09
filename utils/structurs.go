package utils

import "database/sql"

type Handledb struct {
	DB *sql.DB
}

type Post struct {
	Id       int
	UserName string
	Title    string
	Content  string
}
