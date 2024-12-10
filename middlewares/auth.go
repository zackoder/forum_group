package middlewares

import (
	"fmt"
	"net/http"
)



func Permission(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var cookie = struct{
			user_id *http.Cookie
			token *http.Cookie
		}{}
		var err error
		cookie.user_id, err = r.Cookie("user_id")
		if err != nil {
			fmt.Printf("%v",err.Error())
			return
		}
		cookie.token , err = r.Cookie("user_token")
		if err != nil {
			fmt.Printf("%v",err.Error())
			return
		}
		fmt.Printf("user_id: %v,\ntoken: %v\n", cookie.user_id.Value, cookie.token.Value)
		next(w, r)
	}
}