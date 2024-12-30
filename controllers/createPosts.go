package controllers

import (
	"net/http"

	"forum/utils"
)

func CreatedPosts(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/profile.html"}
	utils.ExecuteTemplate(w, pages, nil)
}

func Categories(w http.ResponseWriter, r *http.Request) {
	// categoriesName := r.PathValue("nameCategory")
	// fmt.Println(categoriesName)
	// category := TakeCategories(categoriesName)
	// if category < 1 {
	// 	utils.ErrorHandler(w, http.StatusNotFound, "Page not Found", "The page you are looking for is not available!", nil)
	// 	return
	// }
	pages := []string{"views/pages/categories.html"}
	utils.ExecuteTemplate(w, pages, nil)
}
