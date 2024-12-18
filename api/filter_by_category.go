package api

import (
	"encoding/json"
	"fmt"
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
	id := r.PathValue("CategoryId")
	category_id, num_err := strconv.Atoi(id)
	if num_err != nil {
		fmt.Println(num_err.Error())
		return
	}
	var posts []utils.PostsResult
	query := `
		SELECT p.id,p.user_id,p.title,p.content,p.categories,p.date, u.username
		FROM posts p JOIN posts_categories pc
		ON p.id = pc.post_id AND pc.category_id = ? 
		JOIN users u ON u.id = p.user_id;
	`
	rows, err := utils.DB.Query(query, category_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var user_id int
		var categories string
		var p utils.PostsResult // this is a post im use p for short var
		if err = rows.Scan(&p.Id, &user_id, &p.Title, &p.Content, &categories, &p.Date, &p.UserName); err != nil {
			fmt.Println(err.Error())
			return
		}
		p.Categories = strings.Split(categories, ",")
		posts = append(posts, p)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
