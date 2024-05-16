package controllers

import (
	"net/http"
	"rapikan/configs"
	"rapikan/models"

	"github.com/labstack/echo/v4"
)

func GetServices(c echo.Context) error {
	var services []models.Services

	if err := configs.DB.Preload("Housekeepers").Find(&services).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Services",
		"data":    services,
	})
}

func CreateService(c echo.Context) error {
	service := models.Services{}

	c.Bind(&service)

	// check if service name already exists
	var existingService models.Services
	if err := configs.DB.Where("name = ?", service.Name).First(&existingService).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Service name already exists",
		})
	}

	// check if price is negative
	if service.Price < 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Price cannot be negative",
		})
	}

	// Check the last ID and add autoincrement
	ServiceIDCount := models.Services{}
	configs.DB.Model(&models.Services{}).Order("id desc").First(&ServiceIDCount)

	service.ID = int(ServiceIDCount.ID + 1)

	if err := configs.DB.Create(&service).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create service",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Service",
		"data":    service,
	})
}
