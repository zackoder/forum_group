package main

import (
	"database/sql"
	"fmt"
	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	RunDatabase()
	fmt.Println("hello forum!")
}

func RunDatabase() {
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	models.InitTables(db)
}
