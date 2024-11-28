package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func ExecuteTemplate(w http.ResponseWriter, pages []string, data any) {
	pages = append(pages, "views/base.html")
	tmpl, err := template.ParseFiles(pages...)
	if err != nil {
		fmt.Printf("html page template error: %v", err.Error())
		return
	}
	
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Printf("html page template error: %v", err.Error())
		return
	}
}
