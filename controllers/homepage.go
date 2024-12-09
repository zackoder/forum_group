package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/utils"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusMethodNotAllowed)
		return
	}
	// tmpl, err := template.ParseFiles("../views/home.html") // path page html
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	var isConnected bool
	cookie, err := r.Cookie("token") // Name the cookie
	if err != nil {
		isConnected = false
	} else {
		isConnected = true
		name := ""
		var userId int
		if err := utils.DB.QueryRow(`SELECT user_id FROM sessions WHERE token =?`, cookie.Value).Scan(&userId); err != nil {
			isConnected = false
		}
		if err := utils.DB.QueryRow(`SELECT username FROM users WHERE id =?`, userId).Scan(&name); err != nil {
			// http.Error(w, "Not Found this User", http.StatusInternalServerError)
			fmt.Println("Not Found user Name")
			isConnected = false
		}

	}

	// query := `SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT 20`
	query := "SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT 20"

	rows, err := utils.DB.Query(query)
	if err != nil {
		http.Error(w, "Post Not Found", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var Posts []utils.Post
	for rows.Next() {
		var post utils.Post
		var userId int
		if err := rows.Scan(&post.Id, &userId, &post.Title, &post.Content); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if err := utils.DB.QueryRow("SELECT username FROM users WHERE id =?", userId).Scan(&post.Username); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		Posts = append(Posts, post)

	}
	if isConnected {
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Posts) // create struct data
	// tmpl.Execute(w, Posts)
}
