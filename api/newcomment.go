package api

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func NewComment(w http.ResponseWriter, r *http.Request) {
	var comment utils.Comment
	if r.Method != "POST" {
		fmt.Println("error method not allowd!")
	}
	token, tokenErr := r.Cookie("token")
	if tokenErr != nil {
		fmt.Println(tokenErr.Error())
		return
	}
	postId := r.PathValue("PostId")
	var postIdErr error
	comment.PostId, postIdErr = strconv.Atoi(postId)
	if postIdErr != nil || CheckPost(comment.PostId) {
		fmt.Println("Post id errrrrrrrrrrrrrrrrrrrrrrrrrr")
		fmt.Println(comment.PostId)
		return
	}
	var user_err error
	comment.UserId, user_err = utils.GetUserByToken(token.Value)
	if user_err != nil {
		fmt.Println("get user_id from db errrrrrrrrrrrrrrrrrrrrrr")
		fmt.Println(user_err.Error())
		return
	}
	comment.Comment = r.FormValue("comment")
	query := `INSERT INTO comments(user_id,post_id,comment) VALUES (?,?,?);`
	_, err := utils.DB.Exec(query, comment.UserId, comment.PostId, comment.Comment)
	if err != nil {
		fmt.Println("create comment errrrrrrrrrrrrrrrrrrrrrr")
		fmt.Println(err.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CheckPost(postId int) bool {
	query := `SELECT id FROM posts WHERE id=?;`
	var id int
	err := utils.DB.QueryRow(query, postId).Scan(&id)
	return err != nil
}
