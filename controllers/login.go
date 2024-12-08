package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"forum/utils"

	"github.com/gofrs/uuid/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/login.html"}
	utils.ExecuteTemplate(w, pages, nil)
}

func (db *Date) SingIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	userInf := r.FormValue("userInf")
	userInf = strings.TrimLeft(userInf, " ")
	passwd := r.FormValue("passwd")
	if !IsValidEmail(userInf) && !IsValidUsername(userInf) || passwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "check you input")
		return
	}
	id, err := db.Select(userInf, passwd)
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
	err = db.CraeteSession(id, uid.String())
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
