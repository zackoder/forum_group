package controllers

import (
	"encoding/json"
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
	if strings.Contains(err.Error(), "email") {
		return http.StatusFound, -1, fmt.Errorf("email already used try anther email")
	}
	return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := utils.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if !IsValidUsername(user.Username) || !IsValidEmail(user.Email) || user.Password != user.ConfPass || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Inavlid input for login"})
		return
	}
	if !isValidPassword(user.Password) {
		
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "password is weak"})
		return
	}
	user.Password, err = HasPassowd(user.Password)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})

		return
	}
	statuscode, userId, err := Insert(user)
	if err != nil {
		w.WriteHeader(statuscode)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	uid, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
	}
	err = CraeteSession(userId, uid.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    uid.String(),
		MaxAge:   300,
		HttpOnly: true,
		Path:     "/",
	})
}
