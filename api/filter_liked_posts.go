package api

import (
	"forum/utils"
	"net/http"
)

/*
	- check user authentication
	- get user_id by token
	- get liked posts by user_id
*/
func LikedPosts(w http.ResponseWriter,r *http.Request) {
	token, token_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: token_err,Code: http.StatusInternalServerError},w) {
		return
	}
	var user_id int
	/* i'm not prepare this query because the tocken is verified in middleware */
	get_user_id := `SELECT user_id FROM sessions WHERE token = ? LIMIT 1;`
	err := utils.DB.QueryRow(get_user_id,token.Value).Scan(&user_id)
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError},w) {
		return
	}

}
