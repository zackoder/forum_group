package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"forum/api"
	"forum/controllers"
	"forum/middleware"
	"forum/models"
	"forum/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/* port handling */
	port := ":8001" // use env variable

	/* init database tables */
	var err error
	utils.DB, err = sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer utils.DB.Close()
	models.InitTables(utils.DB)

	/* server mux router */
	mux := http.NewServeMux()
	mux.HandleFunc("/static/", controllers.Server)
	/* pages handlers */
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/add-post", middleware.Authorization(controllers.AddPost))

	/*  */
	mux.HandleFunc("/user/singup", controllers.SingIn)
	mux.HandleFunc("/createpost", controllers.CreatePost)
	mux.HandleFunc("/Register", controllers.RegisterUser)
	mux.HandleFunc("/Login", controllers.SingIn)

	/* api handlers */
	mux.HandleFunc(`/api/{PostId}/comments`, api.Comments)
	mux.HandleFunc("/api/posts", api.FetchPosts)
	mux.HandleFunc("/api/{PostId}/comment/new", middleware.Authorization(api.NewComment))
	mux.HandleFunc("/api/comment/reaction/{PostId}", api.CommentReaction)

	/* filters */
	mux.HandleFunc("/api/category/filter/{CategoryId}", api.FilterByCategory)
	mux.HandleFunc("/api/created/posts", api.CreatedPosts)
	mux.HandleFunc("/api/liked/posts", api.LikedPosts)

	
	/* run server */
	fmt.Printf("server running on http://localhost%s\n", port)
	server_err := http.ListenAndServe(port, mux)
	if server_err != nil {
		fmt.Printf("server runnig error! %v", server_err.Error())
		os.Exit(1)
	}
}
