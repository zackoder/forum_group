package utils

import (
	"log"
	"net/http"
)

func HandleError(err Error, w http.ResponseWriter) bool {
	if err.Err != nil {
		http.Error(w, http.StatusText(err.Code), err.Code)
		log.Println(err.Err)
		return true
	}
	return false
}
