package controllers

import "gorm.io/gorm"

type controller struct {
	db *gorm.DB
}

func New(db *gorm.DB) controller {
	return controller{db: db}
}
