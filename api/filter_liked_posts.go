package api

import (
	"net/http"

	"forum/utils"
)

/*
- check user authentication
- get user_id by token
- get liked posts by user_id
*/
func LikedPosts(w http.ResponseWriter, r *http.Request) {
	// var posts utils.PostsResult
	token, token_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: token_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	var user_id int
	/* i'm not prepare this query because the tocken is verified in middleware */
	get_user_id := `SELECT user_id FROM sessions WHERE token = ? LIMIT 1;`
	err := utils.DB.QueryRow(get_user_id, token.Value).Scan(&user_id)
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
		return
	}
	// get_posts := `SELECT p.id,p.title,p.content,p.categories FROM posts p JOIN users u LEFT JOIN reactions r ON r.user_id = ?`
}

/*
type PostsResult struct {
	Id         int
	UserName   string
	Title      string
	Content    string
	Categories []string
	Date       string
	Reactions  struct {
		Likes    int
		Dislikes int
		Action   string
	}
}
*/
