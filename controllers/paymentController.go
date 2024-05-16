package controllers

import (
	"net/http"
	"rapikan/configs"
	"rapikan/models"
	"time"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) error {
	var payments []models.Payments

	if err := configs.DB.Preload("Bookings").Find(&payments).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Payments",
		"data":    payments,
	})
}

func CreatePayment(c echo.Context) error {
	payment := models.Payments{}
	c.Bind(&payment)

	// Id autoincrement
	PaymentIDCount := models.Payments{}
	configs.DB.Model(&models.Payments{}).Order("id desc").First(&PaymentIDCount)
	payment.ID = int(PaymentIDCount.ID + 1)

	// Set the created_at and updated_at
	payment.Created_At = time.Now()
	payment.Updated_At = time.Now()

	if err := configs.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to create payment",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create payment",
		"data":    payment,
	})
	// return c.JSON(http.StatusOK, payment)
}
