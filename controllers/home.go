package controllers

import (
	"net/http"

	"forum/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ExecuteTemplate(w, []string{"views/pages/error.html"}, nil)
		return
	}
	
	pages := []string{"views/pages/home.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
