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
	webMux.HandleFunc("/", m.CheckMethod(controllers.Home,"GET"))
	webMux.HandleFunc("/add-post", m.CheckMethod(m.Authorization(controllers.CreatePost),"POST"))
	/* --------------------------- login and register handlers --------------------------- */
	webMux.HandleFunc("/user/register", m.CheckMethod(auth.SingUp,"POST"))
	webMux.HandleFunc("/user/login", m.CheckMethod(auth.SingIn,"POST"))
	webMux.HandleFunc("/register", m.CheckMethod(controllers.Register,"GET"))
	webMux.HandleFunc("/login", m.CheckMethod(controllers.Login,"GET"))
	webMux.HandleFunc("/logout", m.Authorization(auth.Logout))
	webMux.HandleFunc("/liked-post", m.CheckMethod(m.Authorization(controllers.LikedPostsPage),"GET"))
	webMux.HandleFunc("/profile", m.CheckMethod(m.Authorization(controllers.CreatedPosts),"GET"))
	webMux.HandleFunc("/categories/{nameCategory}", m.CheckMethod(controllers.Categories,"GET"))

	return webMux
}
