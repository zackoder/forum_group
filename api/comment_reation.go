package api

import (
	"errors"
	"net/http"
	"strconv"

	"forum/utils"
)

func CommentReaction(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	cookie, cookie_err := r.Cookie("token")
	if utils.HandleError(utils.Error{W: w, Err: cookie_err, Code: http.StatusUnauthorized}) {
		return
	}
	get_user := `SELECT user_id FROM sessions WHERE token= ? LIMIT 1`
	stm, stm_err := utils.DB.Prepare(get_user)
	if utils.HandleError(utils.Error{W: w, Err: stm_err, Code: http.StatusInternalServerError}) {
		return
	}
	result, data_err := stm.Exec(cookie.Value)
	if utils.HandleError(utils.Error{W: w, Err: data_err, Code: http.StatusInternalServerError}) {
		return
	}
	user_id, err := result.RowsAffected()
	if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
		return
	}
	reactInfo := struct {
		user_id    int
		comment_id int
	}{1, 1}
	reactInfo.comment_id, err = strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusNotFound}) {
		return
	}
	if r.Method == http.MethodPost {
		var exist int
		query := `SELECT EXISTS(SELECT 1 FROM reactions WHERE (user_id = ? AND comment_id = ?));`
		stm, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
			return
		}
		row_err := stm.QueryRow(user_id, reactInfo.comment_id).Scan(&exist)
		if utils.HandleError(utils.Error{W: w, Err: row_err, Code: http.StatusInternalServerError}) {
			return
		}
		if exist == 0 { /* add reaction */
			query = `INSERT INTO reactions(user_id,comment_id,type) VALUES (?,?,?)`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
				return
			}
			_, err = stm.Exec(&reactInfo.user_id, &reactInfo.comment_id, &action)
			if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
				return
			}
		} else { /* edit reaction */
			query = `UPDATE reactions SET type = ? WHERE user_id = ? AND comment_id = ?;`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
				return
			}
			_, row_err = stm.Exec(query, &action, &reactInfo.user_id, &reactInfo.comment_id)
			if utils.HandleError(utils.Error{W: w, Err: row_err, Code: http.StatusInternalServerError}) {
				return
			}
		}
	} else if r.Method == http.MethodDelete {
		/* delete reaction */
		query := `DELETE FROM reactions WHERE user_id = ? AND comment_id = ?;`
		stmt, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
			return
		}
		_, err = stmt.Exec(reactInfo.user_id, reactInfo.comment_id)
		if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusInternalServerError}) {
			return
		}
	} else {
		/* handle error method not allowed */
		err := errors.New("method not allowd")
		if utils.HandleError(utils.Error{W: w, Err: err, Code: http.StatusMethodNotAllowed}) {
			return
		}
	}
}
