package models

import (
	"time"
)

type Housekeepers struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Name       string    `json:"name" `
	Phone_num  string    `json:"phone_num" `
	Password   string    `json:"password" `
	Created_At time.Time `json:"created_at" `
	Updated_At time.Time `json:"updated_at" `

	ServicesID int      `json:"servicesid"`
	Services   Services `gorm:"foreignKey:ServicesID"`

	Schedules []Schedules `gorm:"foreignKey:HousekeepersID"` // one to many
	Bookings  []Bookings  `gorm:"foreignKey:HousekeepersID"` // one to many
}

// Schedules_ID int       `json:"schedule_id"`
// Schedules    Schedules `gorm:"foreignKey:Schedules_ID;references:ID"`
