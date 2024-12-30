package api

import (
	"encoding/json"
	"errors"
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message": http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	var posts []utils.PostsResult
	/* ---------------------------- Handle token ---------------------------- */
	token, token_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: token_err, Code: http.StatusUnauthorized}, w) {
		return
	}

	/* ---------------------------- Get   from session ---------------------------- */
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
		p.Reactions = GetReaction(user_id, p.Id, "post_id")
		p.Categories = strings.Split(categories, ",")
		fmt.Println(p.Categories)
		posts = append(posts, p)
	}

	/* -------------------------- handle error no content -------------------------- */
	if len(posts) == 0 {
		err := errors.New("no posts")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusNoContent}, w) {
			return
		}
	}

	/* -------------------------- Set result in json response -------------------------- */
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
