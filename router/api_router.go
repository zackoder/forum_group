package router

import (
	"net/http"

	"forum/api"
	"forum/middleware"
)

func APIRouter() *http.ServeMux {
	/* --------------------------- api handlers --------------------------- */
	apiMux := http.NewServeMux()
	apiMux.HandleFunc(`/api/{PostId}/comments`, middleware.CheckMethod(api.Comments, "GET"))                    // comments list
	apiMux.HandleFunc("/api/category/list", middleware.CheckMethod(api.CategoryList, "GET"))                    // get all categories
	apiMux.HandleFunc("/api/posts", middleware.CheckMethod(api.FetchPosts, "GET"))                              // get all posts
	apiMux.HandleFunc("/api/category/filter/{CategoryId}", middleware.CheckMethod(api.FilterByCategory, "GET")) // not complated

	apiMux.HandleFunc("/api/{PostId}/comment/new", middleware.Authorization(api.NewComment))              // create comment
	apiMux.HandleFunc("/api/comment/reaction/{CommentId}", middleware.Authorization(api.CommentReaction)) // (like or dislike) a comment
	apiMux.HandleFunc("/api/posts/reaction/{PostId}", middleware.Authorization(api.PostReaction))         // create new post
	apiMux.HandleFunc("/api/created/posts", middleware.Authorization(api.CreatedPosts))                   // not complated
	apiMux.HandleFunc("/api/liked/posts", middleware.Authorization(api.LikedPosts))                       // not complated
	return apiMux
}
