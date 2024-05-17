package configs

import (
	"rapikan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "root:12345678@tcp(localhost:3306)/mini_project1?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:zPHXmQMHmGHoDNXypPcLKWQBPZrsjnqy@tcp(viaduct.proxy.rlwy.net:57758)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Assign the value to the existing DB variable

	if err != nil {
		panic(err.Error())
	}
	InitMigrate()

}

func InitMigrate() {
	DB.AutoMigrate(&models.Payments{})
	DB.AutoMigrate(&models.Schedules{})
	DB.AutoMigrate(&models.Services{})
	DB.AutoMigrate(&models.Bookings{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Housekeepers{})
}
