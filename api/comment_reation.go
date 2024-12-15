package api

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func CommentReaction(w http.ResponseWriter, r *http.Request) {
	var err error
	action := r.FormValue("action")
	cookie, cookie_err := r.Cookie("token")
	if cookie_err != nil {
		fmt.Println(cookie_err.Error(), "cookie err")
		return
	}
	get_user := `SELECT user_id FROM sessions WHERE token= ? LIMIT 1`
	stm, stm_err := utils.DB.Prepare(get_user)
	if stm_err != nil {
		fmt.Println(stm_err.Error())
		return
	}
	result, data_err := stm.Exec(cookie.Value)
	if data_err != nil {
		fmt.Println(data_err.Error())
		return
	}
	user_id, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reactInfo := struct {
		user_id    int
		comment_id int
	}{1, 1}
	reactInfo.comment_id, err = strconv.Atoi(r.PathValue("PostId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method == http.MethodPost {
		var exist int
		query := `SELECT EXISTS(SELECT 1 FROM reactions WHERE (user_id = ? AND comment_id = ?));`
		stm, err := utils.DB.Prepare(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		row_err := stm.QueryRow(user_id, reactInfo.comment_id).Scan(&exist)
		if row_err != nil {
			fmt.Println(row_err, "exist")
			return
		}
		if exist == 0 { /* add reaction */
			query = `INSERT INTO reactions(user_id,comment_id,type) VALUES (?,?,?)`
			stm, err := utils.DB.Prepare(query)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = stm.Exec(&reactInfo.user_id, &reactInfo.comment_id, &action)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else { /* edit reaction */
			query = `UPDATE reactions SET type = ? WHERE user_id = ? AND comment_id = ?;`
			_, row_err = utils.DB.Exec(query, &action, &reactInfo.user_id, &reactInfo.comment_id)
			if row_err != nil {
				fmt.Println(row_err, "update err")
				return
			}
		}
	} else if r.Method == http.MethodDelete {
		/* delete reaction */
		query := `DELETE FROM reactions WHERE user_id = ? AND comment_id = ?;`
		stmt, err := utils.DB.Prepare(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = stmt.Exec(reactInfo.user_id, reactInfo.comment_id)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
