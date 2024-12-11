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
	// user_id, _ := r.Cookie("user_id")
	// user_token, _ := r.Cookie("user_token")

	// fmt.Printf("user id: %s,\nuser token: %s\n", user_id.Value, user_token.Value)
	utils.ExecuteTemplate(w, pages, nil)
}

func SingIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	userInf := r.FormValue("userInf")
	// userInf = strings.TrimLeft(userInf, " ")
	passwd := r.FormValue("passwd")
	if !IsValidEmail(userInf) && !IsValidUsername(userInf) || passwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "check you input")
		return
	}
	id, err := Select(userInf, passwd)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v", err)
		return
	}
	uid, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "there are error in server try later please")
		return
	}
	err = CraeteSession(id, uid.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "there are error in server try later please")
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
	query := `SELECT id , passwd FROM user
		WHERE user_name = ? OR email = ?
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return -2, fmt.Errorf("problem in server try anther time")
	}
	defer stmt.Close()
	var hashpasswd string
	var id int
	err = stmt.QueryRow(userIfo, userIfo).Scan(&id, &hashpasswd)
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
