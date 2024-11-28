package controllers

import (
	"forum/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/home.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
