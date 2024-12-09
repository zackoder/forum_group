package main

import (
	"fmt"
	"net/http"
	"os"

	"forum/api/handlers"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/api/comments", handlers.Comments)
	mux.HandleFunc("/api/posts", handlers.FetchPosts)
	mux.HandleFunc("/api/reaction", handlers.Reactions)

	/* run server */
	fmt.Printf("server running on http://localhost%s\n", port)
	server_err := http.ListenAndServe(port, mux)
	if server_err != nil {
		fmt.Printf("server runnig error! %v", server_err.Error())
		os.Exit(1)
	}
}
