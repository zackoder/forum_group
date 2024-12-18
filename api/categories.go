package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"forum/utils"
)

func CategoryList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New("method not allowd")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusMethodNotAllowed}, w) {
			return
		}
	}
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
