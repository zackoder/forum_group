package api

import "net/http"

func CommentReaction(w http.ResponseWriter,r *http.Request) {
	action := r.FormValue("action")
	if action == "like" {
		
	} else if action == "dislike" {

	}
}