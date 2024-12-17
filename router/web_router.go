package router

import (
	"net/http"

	"forum/controllers"
	"forum/controllers/auth"
	"forum/middleware"
)

func WebRouter() *http.ServeMux {
	webMux := http.NewServeMux()

	webMux.HandleFunc("/static/", controllers.Server)
	/* pages handlers */
	webMux.HandleFunc("/", controllers.Home)
	webMux.HandleFunc("/add-post", middleware.Authorization(controllers.CreatePost))
	/* login and register handlers */
	webMux.HandleFunc("/Register", auth.RegisterUser)
	webMux.HandleFunc("/Login", auth.SingIn)
	webMux.HandleFunc("/register", controllers.Register)
	webMux.HandleFunc("/login", controllers.Login)

	return webMux
}
