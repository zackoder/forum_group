package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/utils"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
		json.NewEncoder(w).Encode(err)
	}
	cookie, err := r.Cookie("token") // Name the Cookie
	if err != nil {
		// http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		err := Error{Message: "Unauthorized", Code: http.StatusUnauthorized}
		json.NewEncoder(w).Encode(err)
		return
	}

	title := r.FormValue("Title")
	content := r.FormValue("Content")
	categories := r.Form["options"]
	fmt.Println(categories)
	// for i := 0; i < len(categories); i++ {
	// 	fmt.Printf(categories[i])
	// }


	var userId int

	err = utils.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", cookie.Value).Scan(&userId)
	if err != nil {
		err := Error{Message: "Error", Code: 500}
		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		json.NewEncoder(w).Encode(err)
		return
	}
	// parsedDate, err := time.Parse("2006-01-02 15:04:05", date)
	// if err != nil {
	// 	log.Fatal("Invalid date formate", err)
	// }
	// formattedDate := parsedDate.Format("2006-01-02 15:04:05")
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)
	if title == "" || content == "" {
		err := Error{Message: "Title or Content is", Code: http.StatusUnauthorized}
		json.NewEncoder(w).Encode(err)
		return
	}
	result, err := utils.DB.Exec("INSERT INTO posts(user_id, title, content, categories) VALUES(?, ?, ?, ?)", userId, title, strings.ReplaceAll(content, "\r\n", "<br>"), strings.Join(categories, ", "))
	if err != nil {
		err := Error{Message: "can insert in base donne", Code: http.StatusUnauthorized}
		json.NewEncoder(w).Encode(err)
		return
	}
	last_post_id, err := result.LastInsertId()
	fmt.Println(last_post_id)
	if err != nil {
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		err := Error{Message: "Error", Code: http.StatusInternalServerError}
		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		json.NewEncoder(w).Encode(err)
		return
	}
	for _, categ := range categories {
		var category_id int
		err := utils.DB.QueryRow("SELECT id FROM categories WHERE name = ?", categ).Scan(&category_id)
		if err != nil {
			err := Error{Message: "Bad Request", Code: http.StatusBadRequest}
			json.NewEncoder(w).Encode(err)
			return
		}
		_, err = utils.DB.Exec("INSERT INTO posts_categories(post_id, category_id) VALUES(?, ?)", last_post_id, category_id) // GetLast id in table posts
		if err != nil {
			// http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			err := Error{Message: "Bad Request", Code: http.StatusInternalServerError}
			json.NewEncoder(w).Encode(err)
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
