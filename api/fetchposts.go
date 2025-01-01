package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

type Error struct {
	Message string
	Code    int
}

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	offset := r.URL.Query().Get("offset")
	nbr_offset, err := strconv.Atoi(offset)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}
	userid := 0
	cookie, err := r.Cookie("token")
	if err == nil {
		userid = TakeuserId(cookie.Value)
	}
	query := "SELECT id, user_id, title, content, categories, date FROM posts ORDER BY id DESC LIMIT ? OFFSET ?"
	rows, err := utils.DB.Query(query, 20, nbr_offset)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	posts := []utils.PostsResult{}
	// rows, err := stm.Query()
	for rows.Next() {
		var post utils.PostsResult
		var user_id int
		var categories string
		if err := rows.Scan(&post.Id, &user_id, &post.Title, &post.Content, &categories, &post.Date); err != nil {
			continue
		}
		post.Categories = strings.Split(categories, ",")
		post.UserName, _ = GetUsername(user_id)
		post.Reactions = GetReaction(userid, post.Id, "post_id")
		posts = append(posts, post)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
