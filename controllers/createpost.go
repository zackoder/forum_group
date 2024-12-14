package controllers

import (
	"forum/utils"
	"net/http"
	"strings"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	cookie, err := r.Cookie("token") // Name the Cookie
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["option"]

	var userId int
	err = utils.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", cookie.Value).Scan(&userId)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	_, err = utils.DB.Exec("INSERT INTO posts(user_id, title, content, image, categories, date) VALUES(?, ?, ?, ?, ?, ?)", userId, title, strings.ReplaceAll(strings.TrimSpace(content), "\r\n", "<br>"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for _, categ := range categories {
		_, err = utils.DB.Exec("INSERT INTO category(name, post_id) VALUES(?, ?)", categ) // GetLast id in table posts
		if err != nil {
			// Handle Error
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

	// tmpl, err := template.ParseFiles("./templates/createpost.html")
	// if err != nil {
	// 	http.Error(w, "Error in the Parse File", http.StatusInternalServerError)
	// }
	// tmpl.Execute(w, nil)
}
