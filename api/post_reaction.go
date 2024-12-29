package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func PostReaction(w http.ResponseWriter, r *http.Request) {
	var reactPost struct {
		user_id int    // get from token
		post_id int    // get from url
		action  string // get from form
	}
	type Result struct {
		Message string
		Code    int
	}

	/* ----------------------------------- Handle Post Id ----------------------------------- */
	var err error
	reactPost.post_id, err = strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) {
		return
	}

	/* ----------------------------------- handle action ----------------------------------- */
	reactPost.action = r.FormValue("action")
	fmt.Println(r.FormValue("action"))
	/* ----------------------------------- Handle User Id ----------------------------------- */
	token, token_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: token_err, Code: http.StatusUnauthorized}, w) {
		return
	}
	get_user := `SELECT user_id FROM sessions WHERE token= ? LIMIT 1`
	stmt, stmt_err := utils.DB.Prepare(get_user)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	stmt_err = stmt.QueryRow(token.Value).Scan(&reactPost.user_id)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}

	/* ------------------------------ handle post reaction ------------------------------ */
	if r.Method == http.MethodPost {
		var exist int
		query := `SELECT EXISTS(SELECT 1 FROM reactions WHERE (user_id = ? AND post_id = ?));`
		stm, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		row_err := stm.QueryRow(reactPost.user_id, reactPost.post_id).Scan(&exist)
		if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
			return
		}
		if exist == 0 { /* add reaction */
			query = `INSERT INTO reactions(user_id,post_id,type) VALUES (?,?,?)`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
			_, err = stm.Exec(&reactPost.user_id, &reactPost.post_id, &reactPost.action)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
			defer stm.Close()
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Result{Message: http.StatusText(http.StatusCreated), Code: http.StatusCreated})
		} else { /* edit reaction */
			query = `UPDATE reactions SET type = ? WHERE user_id = ? AND post_id = ?;`
			stm, err := utils.DB.Prepare(query)
			if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
				return
			}
			defer stm.Close()
			_, row_err = stm.Exec(&reactPost.action, &reactPost.user_id, &reactPost.post_id)
			if utils.HandleError(utils.Error{Err: row_err, Code: http.StatusInternalServerError}, w) {
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Result{Message: http.StatusText(http.StatusOK), Code: http.StatusOK})
		}
	} else if r.Method == http.MethodDelete {
		/* delete reaction */
		query := `DELETE FROM reactions WHERE user_id = ? AND post_id = ?;`
		stmt, err := utils.DB.Prepare(query)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		_, err = stmt.Exec(reactPost.user_id, reactPost.post_id)
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
