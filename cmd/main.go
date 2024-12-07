package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"forum/controllers"
	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/* port handling */
	port := ":8000" // use env variable

	/* init database tables */
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	models.InitTables(db)

	/* server mux router */
	mux := http.NewServeMux()
	/* run static files */
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	/* pages handlers */
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/add-post", controllers.AddPost)

	/* run server */
	fmt.Printf("server running on http://localhost%s\n", port)
	server_err := http.ListenAndServe(port, mux)
	if server_err != nil {
		fmt.Printf("server runnig error! %v", server_err.Error())
		os.Exit(1)
	}
}