package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"forum/utils"
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
	for i := 0; i < len(categories); i++ {
		fmt.Printf(categories[i])
	}
	image := r.FormValue("image")
	date := r.FormValue("date")

	var userId int
	err = utils.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", cookie.Value).Scan(&userId)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	parsedDate, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		log.Fatal("Invalid date formate", err)
	}
	formattedDate := parsedDate.Format("2006-01-02 15:04:05")
	result, err := utils.DB.Exec("INSERT INTO posts(user_id, title, content, image, categories, date) VALUES(?, ?, ?, ?, ?, ?)", userId, title, strings.ReplaceAll(strings.TrimSpace(content), "\r\n", "<br>"), strings.Join(categories, ", "), image, formattedDate)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	last_post_id, err := result.LastInsertId()
	fmt.Println(last_post_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for _, id_categ := range categories {
		_, err = utils.DB.Exec("INSERT INTO posts_categories(post_id, category_id) VALUES(?, ?)", last_post_id, id_categ) // GetLast id in table posts
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

	// tmpl, err := template.ParseFiles("./templates/createpost.html")
	// if err != nil {
	// 	http.Error(w, "Error in the Parse File", http.StatusInternalServerError)
	// }
	// tmpl.Execute(w, nil)
}
