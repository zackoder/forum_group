package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

func FilterByCategory(w http.ResponseWriter, r *http.Request) {
	/* ---------------------------- check if user is logged and get here id ---------------------------- */
	var logged_user int
	token, token_err := r.Cookie("token")
	if token_err == nil {
		query := `SELECT user_id FROM sessions WHERE token = ?;`
		stmt, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		err = stmt.QueryRow(token.Value).Scan(&logged_user)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
	}

	/* ------------------------------- Handle Category Id ------------------------------- */
	id := r.PathValue("CategoryId")
	category_id, num_err := strconv.Atoi(id)
	if utils.HandleError(utils.Error{Err: num_err, Code: http.StatusNotFound}, w) {
		return
	}

	var posts []utils.PostsResult
	query := `
		SELECT p.id,p.user_id,p.title,p.content,p.categories,p.date, u.username
		FROM posts p JOIN posts_categories pc
		ON p.id = pc.post_id AND pc.category_id = ? 
		JOIN users u ON u.id = p.user_id;
	`
	stmt, stmt_err := utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	rows, rows_err := stmt.Query(category_id)
	if utils.HandleError(utils.Error{Err: rows_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	for rows.Next() {
		var user_id int
		var categories string
		var p utils.PostsResult // this is a post i'm use p for short var
		err := rows.Scan(&p.Id, &user_id, &p.Title, &p.Content, &categories, &p.Date, &p.UserName)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		/* i'm not prepare this query because the post_id is not from user input */
		get_likes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "like")`
		err = utils.DB.QueryRow(get_likes, p.Id).Scan(&p.Reactions.NumLike)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		get_dislikes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "dislike")`
		err = utils.DB.QueryRow(get_dislikes, p.Id).Scan(&p.Reactions.NumDisLike)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		if logged_user > 0 {
			get_action := `SELECT type FROM reactions WHERE (user_id = ? AND post_id = ?)`
			err = utils.DB.QueryRow(get_action, logged_user, p.Id).Scan(&p.Reactions.Action)
			if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
		}
		p.Categories = strings.Split(categories, ",")
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
