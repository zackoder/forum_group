package controllers

import (
	"fmt"
	"net/http"

	"forum/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/login.html"}
	user_id, _ := r.Cookie("user_id")
	user_token, _ := r.Cookie("user_token")

	fmt.Printf("user id: %s,\nuser token: %s\n", user_id.Value, user_token.Value)
	utils.ExecuteTemplate(w, pages, nil)
}
