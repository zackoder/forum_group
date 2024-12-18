package api

import (
	"encoding/json"
	"net/http"

	"forum/utils"
)

func CategoryList(w http.ResponseWriter, r *http.Request) {
	/* ----------------------------------- category list ----------------------------------- */
	query := `SELECT id,name FROM categories`
	rows, err := utils.DB.Query(query)
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
		return
	}
	categories := []utils.Category{}
	for rows.Next() {
		category := utils.Category{}
		if err = rows.Scan(&category.Id, &category.Name); utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		categories = append(categories, category)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
