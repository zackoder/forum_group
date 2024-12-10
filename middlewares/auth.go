package middlewares

import (
	"net/http"

	"forum/utils"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		query := `SELECT user_id FROM sessions WHERE token=?;`
		var user_id int
		row_err := utils.DB.QueryRow(query, token.Value).Scan(&user_id)
		if row_err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
