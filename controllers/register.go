package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/base.html")
	if err != nil {
		fmt.Printf("base page template error: %v", err.Error())
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Printf("base page template error: %v", err.Error())
		return
	}
}
