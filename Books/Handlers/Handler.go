package handlers

import "gorm.io/gorm"

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
