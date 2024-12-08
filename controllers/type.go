package controllers

import "database/sql"

type User struct {
	User_name string `json:"user_name"`
	Email     string `json:"email"`
	Passwd    string `json:"passwd"`
}

type Date struct {
	DB *sql.DB
}

type Reaction struct {
	Nblike   int    `json:"nblike"`
	NbDlike  int    `json:"nbDlike"`
	Reaction string `json:"reaction"`
}
