package controllers

import (
	"net/http"

	"forum/utils"
)

type user struct {
	Name string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // handle invalid path (error 404)
		utils.ExecuteTemplate(w, []string{"views/pages/error.html"}, nil)
		return
	}

	u := &user{Name: "zbessado"}

	token_cookie := http.Cookie{
		Name: "token",
		Value: "123456789abcdefghijklmnopqrstuvwxyz",
	}

	http.SetCookie(w,&token_cookie)

	pages := []string{
		"views/pages/home.html",
		"views/components/new_comment.html",
	}
	utils.ExecuteTemplate(w, pages, u)
}

/*
query := `SELECT * FROM categories;`
	rows, err := utils.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var categories []utils.Category
	for rows.Next() {
		var category utils.Category
		if err := rows.Scan(&category.Id,&category.Name); err != nil {
			log.Println("Failed to scan row:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	fmt.Println(categories)
*/
