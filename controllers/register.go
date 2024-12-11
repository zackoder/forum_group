package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"forum/utils"

	"github.com/gofrs/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/register.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, pages, nil)
	} else if r.Method == http.MethodPost {
		user := utils.User{}
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		user.Username = r.FormValue("username")
		confPass := r.FormValue("password_config")
		if user.Email == "" || user.Password == "" || user.Username == "" || user.Password != confPass {
			fmt.Println(user.Email == "", user.Password == "", user.Username == "", user.Password != confPass)
			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, pages, utils.Error{
				ErrorMs: "invalid Input for register",
			})
			return
		}
		if !IsValidUsername(user.Username) || !IsValidEmail(user.Email) {
			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, pages, utils.Error{
				ErrorMs: "invalid Input for register",
			})
			fmt.Println("ok 1")
			return
		}
		fmt.Println("here")
		var err error
		user.Password, err = HasPassowd(user.Password)
		fmt.Println("ok")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		statuscode, userId, err := Insert(user)
		if err != nil {
			fmt.Println(err)
			if statuscode == http.StatusInternalServerError {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(statuscode)
			utils.ExecuteTemplate(w, pages, utils.Error{
				ErrorMs: err.Error(),
			})
			return
		}
		uid, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = CraeteSession(userId, uid.String())
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    uid.String(),
			MaxAge:   300,
			HttpOnly: true,
			Path:     "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Insert(user utils.User) (int, int, error) {
	query := `INSERT INTO users (username , email , password) 
		VALUES (?, ? , ?)
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err == nil {
		latdID, _ := res.LastInsertId()
		return http.StatusOK, int(latdID), nil
	}
	if strings.Contains(err.Error(), "user_name") {
		return http.StatusFound, -1, fmt.Errorf("user name already used try anther user name")
	} else if strings.Contains(err.Error(), "email") {
		return http.StatusFound, -1, fmt.Errorf("email already used try anther email")
	}
	return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
}
