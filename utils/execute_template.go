package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func ExecuteTemplate(w http.ResponseWriter, pages []string, data any) {
	base_pages := []string{
		"views/components/navbar.html",
		"views/components/footer.html",
		"views/base.html",
	}
	pages = append(pages, base_pages...)
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
