package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName   string `json:"bookname"`
	AuthorName string `json:"authorname"`
	Price      int    `json:"price"`
}

func (book *Book) CheckIsValid() bool {
	if book.BookName != "" && book.AuthorName != "" {
		return true
	}
	return false
}
