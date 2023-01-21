package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "Books.com/charan/books/Models"
	"github.com/gorilla/mux"
)

func (h *handler) UpdataBook(w http.ResponseWriter, r *http.Request) {
	var NewbookData, OldBook models.Book
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&NewbookData)
	r.ParseForm()
	if bookid, hasBookId := params["id"]; hasBookId {
		id, _ := strconv.Atoi(bookid)
		h.db.First(&OldBook, id)
		OldBook.AuthorName = NewbookData.AuthorName
		OldBook.BookName = NewbookData.BookName
		h.db.Save(OldBook)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "updated successfully", "data": OldBook})
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "update failed need to pass proper id"})
	}
}
