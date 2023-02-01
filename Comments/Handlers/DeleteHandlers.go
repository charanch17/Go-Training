package handlers

import (
	"encoding/json"
	"net/http"

	models "comments.com/charan/comments/Models"
	utils "comments.com/charan/comments/Utils"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)
	if comment.CommentId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "requires commentid"})
		return
	}
	data := utils.DeleteCommentRecord(comment.CommentId)
	if data["status"] == "400" {
		w.WriteHeader(http.StatusBadRequest)

	}
	json.NewEncoder(w).Encode(data)

}
