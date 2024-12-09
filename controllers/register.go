package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/register.html"}
	utils.ExecuteTemplate(w, pages, nil)
}

func SingUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	} else if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allown"})
		return
	}
	user := utils.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid input for logup"})
		return
	} else if user.Email == "" || user.Password == "" || user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid input for logup"})
		return
	}
	if !IsValidUsername(user.Username) || !IsValidEmail(user.Email) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "check you input , invalid input"})
		return
	}
	user.Password, err = HasPassowd(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "sorry but there are error in server try anther time"})
		return
	}
	statuscode, err := Insert(user)
	if err != nil {
		w.WriteHeader(statuscode)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "user insert into database"})
}

func Insert(user utils.User) (int, error) {
	query := `INSERT INTO user (user_name , email , passwd) 
		VALUES (?, ? , ?)
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("sorry but there are error in server try anther time")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err == nil {
		return http.StatusOK, nil
	}
	if strings.Contains(err.Error(), "user_name") {
		return http.StatusFound, fmt.Errorf("user name already used try anther user name")
	} else if strings.Contains(err.Error(), "email") {
		return http.StatusFound, fmt.Errorf("email already used try anther email")
	}
	return http.StatusInternalServerError, fmt.Errorf("sorry but there are error in server try anther time")
}
