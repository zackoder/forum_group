package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

type CommentType struct {
	Id        int
	Username  string
	UserImage string
	Comment   string
	Date      string
	Reaction  struct {
		Likes    int
		Dislikes int
		Action   string
	}
}

func Comments(w http.ResponseWriter, r *http.Request) {
	var comments []CommentType
	if r.Method != "GET" {
		fmt.Println("status 405 method not allowd!")
		return
	}
	post_id, err := strconv.Atoi(r.PathValue("PostId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	query := `SELECT id,user_id,comment,date FROM comments WHERE post_id = ?;`
	rows, err := utils.DB.Query(query, post_id)
	if err != nil {
		fmt.Println("query error!")
		return
	}
	for rows.Next() {
		var comment CommentType
		var user_id int
		if cm_err := rows.Scan(&comment.Id, &user_id, &comment.Comment, &comment.Date); cm_err != nil {
			fmt.Println("data parse error!", cm_err.Error())
			return
		}
		comment.Username = GetUsername(user_id)
		comments = append(comments, comment)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func GetUsername(id int) string {
	var username string
	query := `SELECT (username) FROM users WHERE id= ?`
	err := utils.DB.QueryRow(query, id).Scan(&username)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return username
}
