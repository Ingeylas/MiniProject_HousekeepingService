package models

import (
	"time"
)

type Schedules struct {
	ID           int       `json:"id" gorm:"primary_key"`
	Date         time.Time `json:"date"`
	Working_Hour string    `json:"working_hour" `
	Availbility  bool      `json:"availbility"`

	HousekeepersID int          `json:"housekeepersid"`
	Housekeepers   Housekeepers `gorm:"foreignKey:HousekeepersID;references:ID"`

	Bookings []Bookings `gorm:"foreignKey:SchedulesID"` // one to many

	// Housekeepers_ID int `json:"housekeeper_id" gorm:"references:ID"`
}
