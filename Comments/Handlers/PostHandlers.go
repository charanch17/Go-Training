package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	models "comments.com/charan/comments/Models"
	utils "comments.com/charan/comments/Utils"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	uniqId := rand.Intn(1000)
	var data models.Comment
	json.NewDecoder(r.Body).Decode(&data)
	data.CommentId = uniqId
	utils.AppendComment(data)
	json.NewEncoder(w).Encode(data)
}
