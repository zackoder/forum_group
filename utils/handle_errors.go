package utils

import (
	"log"
	"net/http"
)

func HandleError(err Error) bool {
	if err.Err != nil {
		http.Error(err.W, http.StatusText(err.Code), err.Code)
		log.Println(err.Err)
		return true
	}
	return false
}
