package router

import (
	"net/http"

	"forum/api"
	"forum/middleware"
)

func APIRouter() *http.ServeMux {
	/* --------------------------- api handlers --------------------------- */
	apiMux := http.NewServeMux()
	apiMux.HandleFunc(`/api/{PostId}/comments`, api.Comments)                                             // comments list
	apiMux.HandleFunc("/api/{PostId}/comment/new", middleware.Authorization(api.NewComment))              // create comment
	apiMux.HandleFunc("/api/comment/reaction/{CommentId}", middleware.Authorization(api.CommentReaction)) // (like or dislike) a comment
	apiMux.HandleFunc("/api/category/list", api.CategoryList)                                             // get all categories
	apiMux.HandleFunc("/api/posts", api.FetchPosts)
	apiMux.HandleFunc("/api/posts/reaction/{PostId}", middleware.Authorization(api.PostReaction))
	/* filters */
	apiMux.HandleFunc("/api/category/filter/{CategoryId}", api.FilterByCategory)        // not complated
	apiMux.HandleFunc("/api/created/posts", middleware.Authorization(api.CreatedPosts)) // not complated
	apiMux.HandleFunc("/api/liked/posts", middleware.Authorization(api.LikedPosts))     // not complated
	return apiMux
}
