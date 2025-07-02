package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `json:"-"`
	AccountID  uint   `json:"-"`
	Number     string `json:"number"`
	ExpireDate string `json:"expire_date"`
	CVV        string `json:"cvv"`
	IssuedAt   string `json:"issued_at"`
	Status     string `json:"status"`
}
