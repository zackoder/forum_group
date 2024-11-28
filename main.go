package main

import (
	"database/sql"
	"fmt"
	"forum/controllers"
	"forum/models"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/* port handling */
	args := os.Args[1:]
	port := ":8000"
	if len(args) == 1 {
		port = fmt.Sprintf(":%s", port)
	} else if len(args) > 1 {
		fmt.Println("server runnig error You need enter only 1 argument!")
		os.Exit(1)
	}

	/* init database tables */
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	if len(args) > 0 && args[0] == "migrate" {
		models.InitTables(db)
		fmt.Println("database updated successfully :)")
		os.Exit(0)
	}

	/* router */
	mux := http.NewServeMux()

	/* run static files */
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	/* pages handlers */
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)

	/* run server */
	fmt.Printf("server running on http://localhost%s\n", port)
	server_err := http.ListenAndServe(port, mux)
	if server_err != nil {
		fmt.Printf("server runnig error! %v", server_err.Error())
		os.Exit(1)
	}
}
