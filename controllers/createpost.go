package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/api"
	"forum/utils"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(err)
		return
	}
	post := utils.PostsResult{}
	cookie, err := r.Cookie("token") // Name the Cookie
	if err != nil {
		err := Error{Message: "Unauthorized", Code: http.StatusUnauthorized}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	post.Title = r.FormValue("Title")
	post.Content = r.FormValue("Content")
	post.Categories = strings.Split(r.FormValue("options"), ",")
	var userId int

	err = utils.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", cookie.Value).Scan(&userId)
	if err != nil {
		err := Error{Message: "Error", Code: 500}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	if post.Title == "" || post.Content == "" {
		err := Error{Message: "Title or Content is empty", Code: http.StatusBadRequest}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	if len(post.Title) > 100 || len(post.Content) > 1000 {
		err := Error{Message: "Title or Content is more than expact", Code: http.StatusBadRequest}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	result, err := utils.DB.Exec("INSERT INTO posts(user_id, title, content, categories) VALUES(?, ?, ?, ?)", userId, post.Title, post.Content, strings.Join(post.Categories, ","))
	if err != nil {
		err := Error{Message: "can insert in base donne", Code: http.StatusUnauthorized}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	last_post_id, err := result.LastInsertId()
	if err != nil {
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		err := Error{Message: "Error", Code: http.StatusInternalServerError}
		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	post.Id = int(last_post_id)
	for _, categ := range post.Categories {
		category_id := api.TakeCategories(categ)
		if category_id < 1 {
			fmt.Println(err, categ)
			err := Error{Message: "Bad Request", Code: http.StatusBadRequest}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		_, err = utils.DB.Exec("INSERT INTO posts_categories(post_id, category_id) VALUES(?, ?)", post.Id, category_id) // GetLast id in table posts
		if err != nil {
			err := Error{Message: "Bad Request", Code: http.StatusInternalServerError}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	}
	json.NewEncoder(w).Encode(post)
}
