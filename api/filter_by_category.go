package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

type FilterPostsCategory struct {
	UserId     int
	Title      string
	Content    string
	Image      string
	Categories string
	Date       string
}

func FilterByCategory(w http.ResponseWriter, r *http.Request) {
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
		var p utils.PostsResult // this is a post im use p for short var
		err := rows.Scan(&p.Id, &user_id, &p.Title, &p.Content, &categories, &p.Date, &p.UserName)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		p.Categories = strings.Split(categories, ",")
		posts = append(posts, p)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
