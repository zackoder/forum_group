package api

import (
	"encoding/json"
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
	var err error
	reactPost.post_id, err = strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) || CheckPost(reactPost.post_id) {
		return
	}

	/* ----------------------------------- handle action ----------------------------------- */
	reactPost.action = r.FormValue("action")
	// fmt.Println(r.FormValue("action"))
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
	like, err := CheckLIke(reactPost.post_id, reactPost.user_id, "like", "post_id")
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "error in server"})
		return
	}
	dilike, err := CheckLIke(reactPost.post_id, reactPost.user_id, "dislike", "post_id")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "error in server"})
		return
	}
	if reactPost.action == "like" {
		if dilike {
			UpdateLike(reactPost.post_id, reactPost.user_id, "post_id", "like")
		} else if like {
			DeletLike(reactPost.post_id, reactPost.user_id, "post_id")
		} else {
			InsertLike(reactPost.post_id, reactPost.user_id, "post_id", "like")
		}
	} else if reactPost.action == "dislike" {
		if like {
			UpdateLike(reactPost.post_id, reactPost.user_id, "post_id", "dislike")
		} else if dilike {
			DeletLike(reactPost.post_id, reactPost.user_id, "post_id")
		} else {
			InsertLike(reactPost.post_id, reactPost.user_id, "post_id", "dislike")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid format"})
		return
	}
}

func UpdateLike(id, userid int, which, typ string) {
	query := fmt.Sprintf(`
		UPDATE  reactions
		SET type = ?
		WHERE %s = ?
		AND user_id = ?
	`, which)
	_, err := utils.DB.Exec(query, typ, id, userid)
	fmt.Println(err, "update")
}

func InsertLike(id, userid int, which, typ string) {
	query := fmt.Sprintf(`
		INSERT INTO reactions(%s ,user_id ,type)
		VALUES( ? , ? , ? )
	`, which)
	_, err := utils.DB.Exec(query, id, userid, typ)
	fmt.Println(err, "inser")
}

func DeletLike(id, userid int, which string) {
	query := fmt.Sprintf(`
		DELETE FROM reactions WHERE %s = ? AND user_id = ?
	`, which)
	_, err := utils.DB.Exec(query, id, userid)
	fmt.Println(err, "del")
}

func CheckLIke(id, userId int, typ, column string) (bool, error) {
	var like bool
	query := fmt.Sprintf(`
		SELECT EXISTS (
			SELECT 1 
			FROM reactions
			WHERE %s = ? AND 
			user_id = ? AND 
			type = ?
		)
	`, column)
	err := utils.DB.QueryRow(query, id, userId, typ).Scan(&like)
	if err != nil {
		return false, fmt.Errorf("query execution error: %w", err)
	}
	return like, nil
}
