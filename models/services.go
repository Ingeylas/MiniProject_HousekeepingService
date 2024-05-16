package models

type Services struct {
	ID           int            `json:"id" gorm:"primary_key"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Price        float64        `json:"price"`
	Housekeepers []Housekeepers `gorm:"foreignKey:ServicesID"` // one to many
	Bookings     []Bookings     `gorm:"foreignKey:ServicesID"` // one to many
}
