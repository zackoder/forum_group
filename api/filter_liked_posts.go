package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"forum/utils"
)

/*
- check user authentication
- get user_id by token
- get liked posts by user_id
*/
func LikedPosts(w http.ResponseWriter, r *http.Request) {
	var posts []utils.PostsResult
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
	get_posts := `SELECT p.id,p.title,p.content,p.categories,p.date,u.username FROM posts p JOIN users u ON p.user_id = u.id JOIN reactions r ON (r.user_id = ? AND r.post_id = p.id)`
	rows, rows_err := utils.DB.Query(get_posts, user_id)
	if utils.HandleError(utils.Error{Err: rows_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	for rows.Next() {
		var p utils.PostsResult // p is a short name for post
		var categories string
		err := rows.Scan(&p.Id, &p.Title, &p.Content, &categories, &p.Date, &p.UserName)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		/* i'm not prepare this query because the post_id is not from user input */
		get_likes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "like")`
		err = utils.DB.QueryRow(get_likes, p.Id).Scan(&p.Reactions.Likes)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		get_dislikes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "dislike")`
		err = utils.DB.QueryRow(get_dislikes, p.Id).Scan(&p.Reactions.Dislikes)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		get_action := `SELECT type FROM reactions WHERE (user_id = ? AND post_id = ?)`
		err = utils.DB.QueryRow(get_action, user_id, p.Id).Scan(&p.Reactions.Action)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		p.Categories = strings.Split(categories, ",")
		posts = append(posts, p)
	}
	// /* -------------------------- handle error no content -------------------------- */
	// if len(posts) == 0 {
	// 	err := errors.New("no posts")
	// 	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNoContent}, w) {
	// 		return
	// 	}
	// }

	/* -------------------------- Set result in json response -------------------------- */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
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
