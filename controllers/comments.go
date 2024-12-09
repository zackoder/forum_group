package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func Comments(w http.ResponseWriter, r *http.Request) {
	var comment utils.Comment
	if r.Method != "POST" {
		fmt.Println("error method not allowd!")
	}
	user_id, err := r.Cookie("user_id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	comment.UserId, err = strconv.Atoi(user_id.Value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	comment.PostId = 1
	comment.Comment = r.FormValue("comment")
	query := `INSERT INTO comments VALUES (NULL,?,?,?,NULL);`
	_, err = utils.DB.Exec(query, comment.UserId, comment.PostId, comment.Comment)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
