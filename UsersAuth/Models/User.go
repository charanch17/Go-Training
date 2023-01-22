package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

func (u *User) CheckIsValid() bool {
	if u.Email != "" && u.Password != "" {
		return true
	}
	return false
}
