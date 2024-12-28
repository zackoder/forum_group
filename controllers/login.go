package controllers

import (
	"net/http"

	"forum/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/login.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
