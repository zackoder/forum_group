package handlers

import (
	"encoding/json"
	"net/http"
)

type test struct {
	Name string
	City string
}

func Comments(w http.ResponseWriter, r *http.Request) {
	t := &test{Name: "yassin", City: "oujda"}
	json.NewEncoder(w).Encode(t)
}
