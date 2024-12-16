package controllers

import (
	"net/http"

	"forum/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/register.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, pages, nil)
	}
}
