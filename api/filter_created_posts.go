package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/utils"
)

/*
	type PostsResult struct {
	    UserName   string
	    UserImage  string
	    Title      string
	    Content    string
	    Image      string
	    Categories []string
	    Date       string
	    Reactions  struct {
	        Likes    int
	        Dislikes int
	        Action   string
	    }
	}
*/
func CreatedPosts(w http.ResponseWriter, r *http.Request) {
	var createdPosts []utils.PostsResult
	token, token_err := r.Cookie("token")
	if token_err != nil {
		fmt.Println(token_err)
		return
	}
	user_id, user_err := utils.GetUserByToken(token.Value)
	if user_err != nil {
		fmt.Println("get user_id from db errrrrrrrrrrrrrrrrrrrrrr")
		fmt.Println(user_err.Error())
		return
	}
	query := `SELECT u.username,p.title,p.content,p.categories,p.date FROM posts p JOIN users u ON p.user_id = ?`
	rows, rows_err := utils.DB.Query(query, user_id)
	if rows_err != nil {
		fmt.Println(rows_err)
		return
	}
	for rows.Next() {
		var post utils.PostsResult
		var categories string
		if row_err := rows.Scan(&post.UserName, &post.Title, &post.Content, &categories, &post.Date); row_err != nil {
			fmt.Println(row_err)
			return
		}
		post.Categories = strings.Split(categories, ",")
		createdPosts = append(createdPosts, post)
	}

	json.NewEncoder(w).Encode(createdPosts)
}
