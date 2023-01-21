package handlers

import (
	"encoding/json"
	"net/http"

	models "Books.com/charan/books/Models"
)

func (h *handler) AddNewBook(w http.ResponseWriter, r *http.Request) {
	var NewBook models.Book
	json.NewDecoder(r.Body).Decode(&NewBook)
	if NewBook.CheckIsValid() {
		h.db.Create(&NewBook)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": NewBook, "message": "Inserted successfully"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "invalid book data"})
	}

}
