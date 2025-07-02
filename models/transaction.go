package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey"`
	FromAccountID  uint   `json:"-"`
	ToAccountID    uint   `json:"-"`
	SenderUserID   uint   `json:"-"`
	ReceiverUserID uint   `json:"-"`
	Amount         int64  `json:"amount"`
	Type           string `json:"type"`
	Description    string `json:"description,omitempty"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
}
