package models

import "gorm.io/gorm"

type Credit struct {
	gorm.Model
	ID             uint    `gorm:"primaryKey"`
	UserID         uint    `json:"-"`
	AccountID      uint    `json:"-"`
	Amount         int64   `json:"amount"`
	InterestRate   float64 `json:"interest_rate"`
	DurationDays   int     `json:"duration_days"`
	MonthlyPayment int64   `json:"monthly_payment"`
	PaidAmount     int64   `json:"paid_amount" gorm:"default:0"`
	Status         string  `json:"status"`
	IssuedAt       string  `json:"issued_at"`
	DueDate        string  `json:"due_date"`
}
