package middlewares

import (
	"fmt"
	"net/http"
)

func Permission(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			token *http.Cookie
			err   error
		)
		token, err = r.Cookie("user_token")
		if err != nil {
			fmt.Printf("%v", err.Error())
			return
		}
		fmt.Printf("token: %v\n", token.Value)
		next(w, r)
	}
}
