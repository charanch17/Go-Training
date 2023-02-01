package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "comments.com/charan/comments/Models"
	utils "comments.com/charan/comments/Utils"
)

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var updatedData models.Comment
	json.NewDecoder(r.Body).Decode(&updatedData)
	if updatedData.CommentId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "no commentId is passed"})
		return
	}
	comment := utils.SendCommentRef(updatedData.CommentId)
	if comment.CommentId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "no ccomment found with commentid " + strconv.Itoa(updatedData.CommentId)})
		return
	}
	if comment.CommentString != updatedData.CommentString {
		comment.CommentString = updatedData.CommentString
	}
	json.NewEncoder(w).Encode(comment)

}
