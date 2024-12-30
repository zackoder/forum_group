package controllers

import (
	"net/http"

	"forum/utils"
)

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/profile.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, pages, nil)
	}
}

func LikedPosts(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/liked_posts.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, pages, nil)
	}
}
