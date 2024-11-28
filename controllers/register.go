package controllers

import (
	"forum/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/register.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
