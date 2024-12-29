package controllers

import (
	"net/http"

	"forum/utils"
)

type User struct {
	Name string
}

func Home(w http.ResponseWriter, r *http.Request) {
	// userName := ""
	if r.URL.Path != "/" { // handle invalid path (error 404)
		utils.ErrorHandler(w, http.StatusNotFound, "Page not Found", "The page you are looking for is not available!", nil)
		return
	}
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), "this Page doesn't support your Method", nil)
		return
	}
	pages := []string{
		"views/pages/home.html",
		"views/components/new_comment.html",
	}
	utils.ExecuteTemplate(w, pages, nil)
}
