package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"forum/utils"
)

func Comments(w http.ResponseWriter, r *http.Request) {
	/* -------------------------- Get all comments for 1 post with PostId -------------------------- */
	w.Header().Set("Content-type", "application/json")
	var comments []utils.CommentType
	/* -------------------------- Check post_id -------------------------- */
	post_id, err := strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) {
		return
	}
	/* -------------------------- Prepare query -------------------------- */
	query := `SELECT id,user_id,comment,date FROM comments WHERE post_id = ?;`
	stmt, stmt_err := utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	/* -------------------------- Get comments from DB -------------------------- */
	rows, rows_err := stmt.Query(post_id)
	if utils.HandleError(utils.Error{Err: rows_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	for rows.Next() {
		var comment utils.CommentType
		var user_id int
		err := rows.Scan(&comment.Id, &user_id, &comment.Comment, &comment.Date)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		comment.Username, err = GetUsername(user_id, w)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		get_likes := `SELECT COUNT(*) FROM reactions WHERE (comment_id = ? AND type = "like")`
		err = utils.DB.QueryRow(get_likes, comment.Id).Scan(&comment.Reactions.Likes)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		get_dislikes := `SELECT COUNT(*) FROM reactions WHERE (comment_id = ? AND type = "dislike")`
		err = utils.DB.QueryRow(get_dislikes, comment.Id).Scan(&comment.Reactions.Dislikes)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		get_action := `SELECT type FROM reactions WHERE (user_id = ? AND comment_id = ?)`
		err = utils.DB.QueryRow(get_action, user_id, comment.Id).Scan(&comment.Reactions.Action)
		if err != sql.ErrNoRows && utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		comments = append(comments, comment)
	}

	/* -------------------------- handle error no content -------------------------- */


	/* -------------------------- Set result in json response -------------------------- */
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func GetUsername(id int, w http.ResponseWriter) (string, error) {
	/* ----------------------------- this function get username ----------------------------- */
	var username string
	query := `SELECT (username) FROM users WHERE id= ?`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	err = stmt.QueryRow(id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
