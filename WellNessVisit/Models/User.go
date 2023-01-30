package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	PhoneNo   string `json:"phonenumber"`
	RoleID    *uint  `gorm:"not null" json:"roleid"`
	Role      Role
}

func (user *User) CheckIsValid() bool {
	if user.Email == "" || user.Password == "" || len(user.Password) < 8 {
		fmt.Println(user)
		fmt.Println(user.RoleID)
		fmt.Println(user.Role.RoleName)
		return false
	}
	return true
}
