package controllers

import (
	"net/http"

	"forum/utils"
)

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/CreatePost.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, pages, nil)
	}
}
