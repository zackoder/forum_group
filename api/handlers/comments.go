package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/utils"
)

func Comments(w http.ResponseWriter, r *http.Request) {
	var comments []utils.Comment
	if r.Method != "GET" {
		fmt.Println("status 405 method not allowd!")
		return
	}
	query := `SELECT comment FROM comments;`
	rows, err := utils.DB.Query(query)
	if err != nil {
		fmt.Println("query error!")
		return
	}
	for rows.Next() {
		var comment utils.Comment
		if cm_err := rows.Scan(&comment.Comment); cm_err != nil {
			fmt.Println("data parse error!")
			return
		}
		comments = append(comments, comment)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
