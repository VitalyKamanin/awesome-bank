package models

import "gorm.io/gorm"

type PaymentSchedule struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	CreditID   uint   `json:"-"`
	AmountDue  int64  `json:"amount_due"`
	DueDate    string `json:"due_date"`
	PaidAmount int64  `json:"paid_amount" gorm:"default:0"`
	Status     string `json:"status"`
	PaidAt     string `json:"paid_at,omitempty"`
}
