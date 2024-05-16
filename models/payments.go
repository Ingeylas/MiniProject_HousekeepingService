package models

import (
	"time"
)

type Method string

const (
	Virtual_Account Method = "Virtual Account"
	EWallet         Method = "E-Wallet"
)

type Payments struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Method     Method    `json:"method" `
	Created_At time.Time `json:"created_at" `
	Updated_At time.Time `json:"updated_at" `

	Bookings []Bookings `gorm:"foreignKey:PaymentsID"` // one to many
}
