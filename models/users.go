package models

import (
	"time"
)

type Users struct {
	ID         int       `json:"id" from:"id" gorm:"primary_key"`
	Username   string    `json:"username" from:"username"`
	Email      string    `json:"email" from:"email"`
	Phone_num  string    `json:"phone_num" from:"phone_num"`
	Password   string    `json:"password" from:"password"`
	Address    string    `json:"address" from:"address"`
	Created_At time.Time `json:"created_at" from:"created_at"`
	Updated_At time.Time `json:"updated_at" from:"updated_at"`

	Bookings []Bookings `gorm:"foreignKey:UsersID"` // one to many
}

type UserResponse struct {
	ID       int    `json:"id" from:"id"`
	Username string `json:"username" from:"username"`
	Password string `json:"password" from:"password"`
	Token    string `json:"token" form:"password"`
}
