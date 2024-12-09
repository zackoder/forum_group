package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"forum/api/handlers"
	"forum/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/api/comments", handlers.Comments)
	mux.HandleFunc("/api/posts", handlers.FetchPosts)
	mux.HandleFunc("/api/reaction", handlers.Reactions)
	var err error

	utils.DB, err = sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	/* run server */
	fmt.Printf("server running on http://localhost%s\n", port)
	server_err := http.ListenAndServe(port, mux)
	if server_err != nil {
		fmt.Printf("server runnig error! %v", server_err.Error())
		os.Exit(1)
	}
}
