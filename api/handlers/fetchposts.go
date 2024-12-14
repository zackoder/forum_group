package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/utils"
)

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	offset := r.URL.Query().Get("offset")
	nbr_offset, err := strconv.Atoi(offset)
	if err != nil {
		nbr_offset = 0
	}

	query := "SELECT id, user_id, title, content, image, categories, date FROM posts ORDER BY id DESC LIMIT ? OFFSET ?"
	rows, err := utils.DB.Query(query, 20, nbr_offset)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	posts := []utils.Post{}
	for rows.Next() {
		// post := Post{}
		var post utils.Post
		var user_id int
		if err := rows.Scan(&post.Id, &user_id, &post.Title, &post.Content, &post.Image, &post.Categories, &post.Date); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if err := utils.DB.QueryRow("SELECT username FROM users WHERE id = ?", user_id).Scan(&post.Username); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
