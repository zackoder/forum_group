package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"forum/utils"
)

/*
	- get token
	- get user_id from session
	- get posts for the (logged user)
*/

func CreatedPosts(w http.ResponseWriter, r *http.Request) {
	var posts []utils.PostsResult
	/* ---------------------------- Handle token ---------------------------- */
	token, token_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: token_err, Code: http.StatusUnauthorized}, w) {
		return
	}

	/* ---------------------------- Get user_id from session ---------------------------- */
	var user_id int
	query := `SELECT user_id FROM sessions WHERE token = ?;`
	stmt, stmt_err := utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	stmt_err = stmt.QueryRow(token.Value).Scan(&user_id)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}

	/* ---------------------------- Get Created Prosts ---------------------------- */
	query = `SELECT p.id,p.title,p.content,p.categories,p.date,u.username FROM posts p JOIN users u ON (p.user_id = ? AND u.id = ?)`
	stmt, stmt_err = utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	rows, rows_err := stmt.Query(user_id, user_id)
	if utils.HandleError(utils.Error{Err: rows_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	for rows.Next() {
		var p utils.PostsResult // post
		var categories string
		row_err := rows.Scan(&p.Id, &p.Title, &p.Content, &categories, &p.Date, &p.UserName)
		if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
			return
		}
		/* i'm not prepare this query because the post_id is not from user input */
		get_likes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "like")`
		err := utils.DB.QueryRow(get_likes, p.Id).Scan(&p.Reactions.Likes)
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
