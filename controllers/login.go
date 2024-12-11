package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/utils"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/login.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, pages, nil)
	} else if r.Method == http.MethodPost {
		userInf := r.FormValue("email")
		passwd := r.FormValue("password")
		if !IsValidEmail(userInf) || passwd == "" {
			w.WriteHeader(http.StatusBadRequest)
			utils.ExecuteTemplate(w, pages, utils.Error{
				ErrorMs: "Check you input",
			})
			return
		}
		id, err := Select(userInf, passwd)
		if err != nil {
			fmt.Println(err)
			if id == -2 {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNotFound)
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
		err = CraeteSession(id, uid.String())
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
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func SingIn(w http.ResponseWriter, r *http.Request) {
}

func Select(userIfo, passwd string) (int, error) {
	query := `SELECT id , password FROM users
		WHERE email = ?
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return -2, fmt.Errorf("problem in server try anther time")
	}
	defer stmt.Close()
	var hashpasswd string
	var id int
	err = stmt.QueryRow(userIfo).Scan(&id, &hashpasswd)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("user or password not correct")
	} else if err != nil {
		return -1, fmt.Errorf("check your input")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashpasswd), []byte(passwd)) != nil {
		return -1, fmt.Errorf("user or password not correct")
	}
	return id, nil
}

func CraeteSession(userid int, session string) error {
	query := `INSERT INTO session(user_id , uid)
		VALUES(?,?)
		ON CONFLICT DO UPDATE SET uid = EXCLUDED.uid , date = CURRENT_TIMESTAMP
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid, session)

	return err
}
