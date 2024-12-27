package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"forum/utils"
)

func CommentReaction(w http.ResponseWriter, r *http.Request) {
	var reactInfo struct {
		user_id    int    // get from token
		comment_id int    // get from url
		action     string // get from form
	}

	type Result struct {
		Message string
		Code    int
	}

	w.Header().Set("Content-Type", "application/json")

	/* ------------------------------ handle user_id ------------------------------ */
	/* get cookie */
	cookie, cookie_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: cookie_err, Code: http.StatusUnauthorized}, w) {
		return
	}

	/* select user_id from database */
	get_user := `SELECT user_id FROM sessions WHERE token= ? LIMIT 1`
	stm, stm_err := utils.DB.Prepare(get_user)
	if utils.HandleError(utils.Error{Err: stm_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	err := stm.QueryRow(cookie.Value).Scan(&reactInfo.user_id)
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
		return
	}

	/* ------------------------------ handle comment_id ------------------------------ */
	reactInfo.comment_id, err = strconv.Atoi(r.PathValue("CommentId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) {
		fmt.Println("post id not valid")
		return
	}

	/* ------------------------------ handle comment reaction ------------------------------ */
	if r.Method == http.MethodPost {
		/* ------------------------------ handle action ------------------------------ */
		reactInfo.action = r.FormValue("action")
		actions := []string{"like", "dislike"}
		if !slices.Contains(actions, reactInfo.action) {
			err := errors.New("invalid action")
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusBadRequest}, w) {
				return
			}
		}
		var exist int
		query := `SELECT EXISTS(SELECT 1 FROM reactions WHERE (user_id = ? AND comment_id = ?));`
		stm, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		row_err := stm.QueryRow(reactInfo.user_id, reactInfo.comment_id).Scan(&exist)
		if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
			return
		}
		if exist == 0 { /* add reaction */
			query = `INSERT INTO reactions(user_id,comment_id,type) VALUES (?,?,?)`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return /* get user_id with token start */
			}
			_, err = stm.Exec(&reactInfo.user_id, &reactInfo.comment_id, &reactInfo.action)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Result{Message: http.StatusText(http.StatusCreated), Code: http.StatusCreated})
		} else { /* edit reaction */
			query = `UPDATE reactions SET type = ? WHERE user_id = ? AND comment_id = ?;`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
			_, row_err = stm.Exec(&reactInfo.action, &reactInfo.user_id, &reactInfo.comment_id)
			if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Result{Message: http.StatusText(http.StatusOK), Code: http.StatusOK})
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Result{Message: http.StatusText(http.StatusOK), Code: http.StatusOK})
	} else {
		/* handle error method not allowed */
		err := errors.New("method not allowd")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusMethodNotAllowed}, w) {
			return
		}
	}
}
