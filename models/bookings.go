package models

import (
	"time"
)

type Status string

const (
	Awaiting_Payment Status = "Awaiting Payments"
	Paid             Status = "Purchase Successful"
	Cleaning         Status = "Cleaning Your Place"
	Done             Status = "Cleaning Done"
)

type Bookings struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Duration   int       `json:"duration" `
	Status     Status    `json:"status" `
	Total      float64   `json:"total" `
	Created_At time.Time `json:"created_at" `
	Updated_At time.Time `json:"updated_at" `

	UsersID int   `json:"userid"`
	Users   Users `gorm:"foreignKey:UsersID;references:ID"`

	ServicesID int      `json:"servicesid"`
	Services   Services `gorm:"foreignKey:ServicesID;references:ID"`

	HousekeepersID int          `json:"housekeepersid"`
	Housekeepers   Housekeepers `gorm:"foreignKey:HousekeepersID;references:ID"`

	SchedulesID int       `json:"schedulesid"`
	Schedules   Schedules `gorm:"foreignKey:SchedulesID;references:ID"`

	PaymentsID int      `json:"paymentsid"`
	Payments   Payments `gorm:"foreignKey:PaymentsID;references:ID"`

	// User_Id        Users        `json:"user_id" from:"user_id"`
	// Service_Id     Services     `json:"service_id" from:"service_id"`
	// Housekeeper_Id Housekeepers `json:"housekeeper_id" from:"housekeeper_id"`
	// Schedule_Id    Schedules    `json:"schedule_id" from:"schedule_id"`
	// Payment_Id     Payments
}
