package controllers

import (
	"database/sql"
	"encoding/json"
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
	}
}

func SingIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	user := utils.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if !IsValidEmail(user.Email) || user.Password == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Inavlid input for login"})
		return
	}
	id, err := Select(user.Email, user.Password)
	if err != nil {
		if id == -2 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "user or password not correct"})
		}

		return
	}
	uid, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		return
	}
	err = CraeteSession(id, uid.String())
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
	query := `INSERT INTO sessions(user_id , token)
		VALUES(?,?)
		ON CONFLICT DO UPDATE SET token = EXCLUDED.token , date = CURRENT_TIMESTAMP
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid, session)

	return err
}
