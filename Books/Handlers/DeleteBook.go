package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "Books.com/charan/books/Models"
	"github.com/gorilla/mux"
)

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if Bookid, hasBookId := params["id"]; hasBookId {
		id, _ := strconv.Atoi(Bookid)
		h.db.First(&book, id)
		if !book.CheckIsValid() {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "No Book Found"})
			return
		}
		h.db.Delete(&book)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Book Deleted successfully"})
	}
}
