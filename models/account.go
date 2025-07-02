package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `json:"-"`
	Balance int64  `json:"balance" gorm:"default:0"`
	Type    string `json:"type"`
}

func (a *Account) UpdateBalance(db *gorm.DB) error {
	return db.Save(a).Error
}
