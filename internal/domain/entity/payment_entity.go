package entity

import (
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusSuccess   PaymentStatus = "SUCCESS"
	PaymentStatusFailed    PaymentStatus = "FAILED"
	PaymentStatusExpired   PaymentStatus = "EXPIRED"
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
)

type Payment struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID       uint   `json:"order_id" gorm:"not null"`
	Method        string `json:"method" gorm:"not null"`
	TransactionID string `json:"transaction_id" gorm:"not null"`
	Order         Order  `json:"order" gorm:"foreignKey:OrderID"`
	gorm.Model
}

func (p *Payment) TableName() string {
	return "payments"
}
