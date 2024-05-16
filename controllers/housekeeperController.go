package controllers

import (
	"net/http"
	"rapikan/configs"
	"rapikan/models"

	"time"

	"github.com/labstack/echo/v4"
)

func GetHousekeepers(c echo.Context) error {
	var housekeepers []models.Housekeepers

	if err := configs.DB.Preload("Services").Find(&housekeepers).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Housekeepers",
		"data":    housekeepers,
	})
}

func GetHousekeeperByID(c echo.Context) error {
	id := c.Param("id")
	var housekeeper models.Housekeepers

	// if err := configs.DB.Find(&housekeeper, "service_id = ?", id).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{
	// 		"message": "Data Not Found",
	// 	})
	// }

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "Success Get Housekeeper",
	// 	"data":    housekeeper,
	// })

	if err := configs.DB.Preload("Services").Find(&housekeeper, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Housekeeper",
		"data":    housekeeper,
	})
}

func CreateHousekeeper(c echo.Context) error {
	housekeeper := models.Housekeepers{}

	c.Bind(&housekeeper)

	// check if housekeeper name already exists
	var existingHousekeeper models.Housekeepers
	if err := configs.DB.Where("name = ?", housekeeper.Name).First(&existingHousekeeper).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Housekeeper name already exists",
		})
	}

	//Set Time Now
	housekeeper.Created_At = time.Now() // set time now
	housekeeper.Updated_At = time.Now() // set time now

	// Check the last ID and add autoincrement
	HousekeeperIDCount := models.Housekeepers{}
	configs.DB.Model(&models.Housekeepers{}).Order("id desc").First(&HousekeeperIDCount)
	housekeeper.ID = int(HousekeeperIDCount.ID + 1)

	// check if service id exists and add to housekeeper
	service := models.Services{}
	if err := configs.DB.Find(&service, "id = ?", housekeeper.ServicesID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Service Not Found",
		})
	}
	housekeeper.Services = service

	if err := configs.DB.Create(&housekeeper).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create housekeeper",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Housekeeper",
		"data":    housekeeper,
	})
}

func GetBookingsByHousekeeperID(c echo.Context) error {
	id := c.Param("id")
	var bookings []models.Bookings

	if err := configs.DB.Preload("Services").Preload("Payments").Find(&bookings, "housekeepers_id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Bookings",
		"data":    bookings,
	})
}
