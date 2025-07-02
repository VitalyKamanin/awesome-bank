package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

func (u *User) Validate() error {
	if u.Username == "" || u.Email == "" || u.Password == "" {
		return fmt.Errorf("missing required fields")
	}
	return nil
}
