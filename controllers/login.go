package controllers

import (
	"net/http"

	"forum/api"
	"forum/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), "this Page doesn't support your Method", nil)
		return
	}
	cookie, err := r.Cookie("token")
	if err == nil {
		userId := api.TakeuserId(cookie.Value)
		if userId > 1 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	pages := []string{"views/pages/login.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
