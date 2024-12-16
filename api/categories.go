package api

import (
	"encoding/json"
	"forum/utils"
	"net/http"
)

func CategoryList(w http.ResponseWriter,r *http.Request) {
	query := `SELECT id,name FROM categories`
	rows, err := utils.DB.Query(query)
	if utils.HandleError(utils.Error{W:w,Err: err,Code: http.StatusInternalServerError}) {
		return
	}
	var categories []utils.Category
	for rows.Next() {
		var category utils.Category
		if err = rows.Scan(&category.Id,&category.Name); utils.HandleError(utils.Error{W:w,Err: err,Code: http.StatusInternalServerError}) {
			return
		}
		categories = append(categories, category)
	}
	json.NewEncoder(w).Encode(categories)
}