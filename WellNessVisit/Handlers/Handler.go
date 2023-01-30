package handlers

import (
	models "Privia.com/charan/WellNessVisits/Models"
	"gorm.io/gorm"
)

type Handler struct {
	db         *gorm.DB
	activeUser *models.User
}

func New(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) SetactiveUser(user *models.User) {
	h.activeUser = user
}
