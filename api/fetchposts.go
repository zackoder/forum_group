package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	query := "SELECT id, user_id, title, content, categories, date FROM posts ORDER BY id DESC LIMIT ? OFFSET ?"
	// stm, err := utils.DB.Prepare(query)
	rows, err := utils.DB.Query(query, 20, nbr_offset)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	posts := []utils.PostsResult{}
	// rows, err := stm.Query()
	for rows.Next() {
		// post := Post{}
		var post utils.PostsResult
		var user_id int
		// var image sql.NullString
		var date sql.NullString
		var categories string

		// type PostsResult struct {
		// 	Id         int
		// 	UserName   string
		// 	Title      string
		// 	Content    string
		// 	Categories []string
		// 	Date       string
		// 	Reactions  struct {
		// 		Likes    int
		// 		Dislikes int
		// 		Action   string
		// 	}
		// }

		if err := rows.Scan(&post.Id, &user_id, &post.Title, &post.Content, &categories, &date); err != nil {
			// log.Printf("Error scanning row %v", err)

			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		post.Categories = strings.Split(categories, ",")

		// if image.Valid {
		// 	post.Image = image.String
		// } else {
		// 	post.Image = ""
		// }
		if date.Valid {
			post.Date = date.String
		} else {
			post.Date = ""
		}

		if err := utils.DB.QueryRow("SELECT username FROM users WHERE id = ?", user_id).Scan(&post.UserName); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// Get Likes the this post
		get_likes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "like");`
		if err := utils.DB.QueryRow(get_likes, post.Id).Scan(&post.Reactions.Likes); err != nil {
			error := Error{Message: http.StatusText(http.StatusInternalServerError), Code: http.StatusInternalServerError}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		get_dislikes := `SELECT COUNT(*) FROM reactions WHERE (post_id = ? AND type = "dislike");`
		if err := utils.DB.QueryRow(get_dislikes, post.Id).Scan(&post.Reactions.Dislikes); err != nil {
			error := Error{Message: http.StatusText(http.StatusInternalServerError), Code: http.StatusInternalServerError}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}

		if user_id > 0 {
			get_action := `SELECT type FROM reactions WHERE (post_id = ? AND user_id = ?);`
			err_action := utils.DB.QueryRow(get_action, post.Id, user_id).Scan(&post.Reactions.Action)
			if err_action != nil {
				// fmt.Println(err_action.Error())
				if err_action == sql.ErrNoRows {
					post.Reactions.Action = ""
				} else {
					fmt.Println("error")
					error := Error{Message: http.StatusText(http.StatusInternalServerError), Code: http.StatusInternalServerError}
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(error)
					return
				}
			}

		}
		posts = append(posts, post)
		fmt.Println(posts)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
