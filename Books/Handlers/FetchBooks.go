package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	models "Books.com/charan/books/Models"
	"github.com/gorilla/mux"
)

func (h *handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	params := mux.Vars(r)
	var booksList []models.Book
	if bookId, hasBookId := params["id"]; hasBookId {
		id, _ := strconv.Atoi(bookId)
		h.db.Find(&booksList, "id = ?", id)
	} else {
		h.db.Find(&booksList)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksList)
}
