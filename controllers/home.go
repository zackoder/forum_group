package controllers

import (
	"net/http"

	"forum/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // handle invalid path (error 404)
		utils.ExecuteTemplate(w, []string{"views/pages/error.html"}, nil)
		return
	}
	pages := []string{"views/pages/home.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
