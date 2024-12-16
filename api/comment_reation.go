package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func CommentReaction(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	cookie, cookie_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: cookie_err, Code: http.StatusUnauthorized}, w) {
		return
	}

	/* get user_id with token start */
	get_user := `SELECT user_id FROM sessions WHERE token= ? LIMIT 1`
	stm, stm_err := utils.DB.Prepare(get_user)
	if utils.HandleError(utils.Error{Err: stm_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	result, data_err := stm.Exec(cookie.Value)
	if utils.HandleError(utils.Error{Err: data_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	user_id, err := result.RowsAffected()
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
		return
	}
	/* get user_id with token end */

	if user_id == 0 {
		// handle token not valid err or expire session
		err := errors.New("token not valid")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusUnauthorized}, w) {
			return
		}
	}
	reactInfo := struct {
		user_id    int
		comment_id int
	}{1, 1}
	reactInfo.comment_id, err = strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) {
		fmt.Println("post id not valid")
		return
	}
	if r.Method == http.MethodPost {
		var exist int
		query := `SELECT EXISTS(SELECT 1 FROM reactions WHERE (user_id = ? AND comment_id = ?));`
		stm, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		row_err := stm.QueryRow(user_id, reactInfo.comment_id).Scan(&exist)
		if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
			return
		}
		if exist == 0 { /* add reaction */
			query = `INSERT INTO reactions(user_id,comment_id,type) VALUES (?,?,?)`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return/* get user_id with token start */
			}
			_, err = stm.Exec(&reactInfo.user_id, &reactInfo.comment_id, &action)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
		} else { /* edit reaction */
			query = `UPDATE reactions SET type = ? WHERE user_id = ? AND comment_id = ?;`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
			_, row_err = stm.Exec(query, &action, &reactInfo.user_id, &reactInfo.comment_id)
			if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
				return
			}
		}
	} else if r.Method == http.MethodDelete {
		/* delete reaction */
		query := `DELETE FROM reactions WHERE user_id = ? AND comment_id = ?;`
		stmt, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		_, err = stmt.Exec(reactInfo.user_id, reactInfo.comment_id)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
	} else {
		/* handle error method not allowed */
		err := errors.New("method not allowd")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusMethodNotAllowed}, w) {
			return
		}
	}
}
