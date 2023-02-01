package handlers

import (
	"encoding/json"
	"net/http"

	models "comments.com/charan/comments/Models"
	utils "comments.com/charan/comments/Utils"
	"github.com/gorilla/mux"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postid, hasPostId := params["postid"]
	commentorid, hasCommentorId := params["commentorid"]
	var data []models.Comment
	if hasPostId {
		data = utils.GetCommentsByField("PostId", postid)

	} else if hasCommentorId {
		data = utils.GetCommentsByField("CommentorId", commentorid)
	} else {
		data = utils.GetAllComments()
	}
	w.Header().Add("Content-type", "Application/json")
	if len(data) > 0 {
		json.NewEncoder(w).Encode(data)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"error": "no comments found"})

}
