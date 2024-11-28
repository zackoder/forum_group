package controllers

import (
	"forum/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/login.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
