package router

import (
	"net/http"

	"forum/controllers"
	"forum/controllers/auth"
	m "forum/middleware"
)

func WebRouter() *http.ServeMux {
	webMux := http.NewServeMux()

	webMux.HandleFunc("/static/", controllers.Server)
	/* --------------------------- pages handlers --------------------------- */
	webMux.HandleFunc("/", controllers.Home)
	webMux.HandleFunc("/add-post", m.Authorization(controllers.CreatePost))
	/* --------------------------- login and register handlers --------------------------- */
	webMux.HandleFunc("/user/register", auth.SingUp)
	webMux.HandleFunc("/user/login", auth.SingIn)
	webMux.HandleFunc("/register", controllers.Register)
	webMux.HandleFunc("/login", controllers.Login)
	webMux.HandleFunc("/logout", m.Authorization(auth.Logout))
	webMux.HandleFunc("/liked-post", m.Authorization(controllers.LikedPostsPage))
	webMux.HandleFunc("/profile", m.Authorization(controllers.CreatedPosts))
	webMux.HandleFunc("/categories/{nameCategory}", (controllers.Categories))

	return webMux
}
