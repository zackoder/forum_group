package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }
	offset := r.URL.Query().Get("offset")
	nbr_offset, err := strconv.Atoi(offset)
	if err != nil {
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		json.NewEncoder(w).Encode(nil)
		return
	}
	//, image, categories, date
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
		var image sql.NullString
		var date sql.NullString
		if err := rows.Scan(&post.Id, &user_id, &post.Title, &post.Content, &image, &post.Categories, &date); err != nil {
			// log.Printf("Error scanning row %v", err)
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if image.Valid {
			post.Image = image.String
		} else {
			post.Image = ""
		}
		if date.Valid {
			post.Date = date.String
		} else {
			post.Date = ""
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
