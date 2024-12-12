package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

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
	var contents []PostsResult
	query := `
		SELECT p.user_id,p.title,p.content,p.categories,p.date, u.username
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
		var c PostsResult
		if err = rows.Scan(&user_id, &c.Title, &c.Content, &categories, &c.Date, &c.UserName); err != nil {
			fmt.Println(err.Error())
			return
		}
		c.Categories = strings.Split(categories, ",")
		contents = append(contents, c)
	}
	json.NewEncoder(w).Encode(contents)
}
