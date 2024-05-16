package controllers

import (
	"net/http"
	"rapikan/configs"
	"rapikan/models"
	"time"

	"github.com/labstack/echo/v4"
)

type BookingData struct {
	Booking models.Bookings `json:"booking"`
	Payment models.Payments `json:"payment"`
}

func GetBookings(c echo.Context) error {
	var bookings []models.Bookings

	if err := configs.DB.Preload("Services").Preload("Payments").Find(&bookings).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Bookings",
		"data":    bookings,
	})
}

// func CreateBooking(c echo.Context) error {
// 	if err := CreatePayment(c); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "Failed to create payment",
// 		})
// 	}
// 	bookingData := BookingData{}

// 	if err := c.Bind(&bookingData); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "Failed to bind booking data",
// 		})
// 	}
// 	booking := models.Bookings{}

// 	booking = bookingData.Booking

// 	// Id autoincrement
// 	BookingIDCount := models.Bookings{}
// 	configs.DB.Model(&models.Bookings{}).Order("id desc").First(&BookingIDCount)
// 	booking.ID = int(BookingIDCount.ID + 1)

// 	// Calculate the total price
// 	// var total float64
// 	total := booking.Services.Price * float64(booking.Duration)
// 	booking.Total = total

// 	// Set the status to Awaiting Payment
// 	booking.Status = models.Awaiting_Payment

// 	// Set the created_at and updated_at
// 	booking.Created_At = time.Now()
// 	booking.Updated_At = time.Now()

// 	if err := configs.DB.Create(&booking).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "Failed to create booking",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Success create booking",
// 		"data":    booking,
// 	})
// }

func CreateBooking(c echo.Context) error {
	bookingData := BookingData{}

	if err := c.Bind(&bookingData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to bind booking data",
		})
	}

	booking := bookingData.Booking
	payment := bookingData.Payment

	// //Check availibility
	// var CheckAvail bool
	// if err := configs.DB.Model(&models.Schedules{}).Select("Availbility").Where("id = ?", booking.SchedulesID).First(&CheckAvail).Error; err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "Failed to check availibility",
	// 	})
	// }
	// if !CheckAvail {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "Housekeeper is not available",
	// 	})
	// }

	//Check availibility
	for i := booking.SchedulesID; i < booking.SchedulesID+booking.Duration; i++ {
		var CheckAvail bool
		if err := configs.DB.Model(&models.Schedules{}).Select("Availbility").Where("id = ?", i).First(&CheckAvail).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Failed to check availibility",
			})
		}
		if !CheckAvail {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Housekeeper is not available",
			})

		}
	}

	// Id autoincrement booking
	BookingIDCount := models.Bookings{}
	configs.DB.Model(&models.Bookings{}).Order("id desc").First(&BookingIDCount)
	booking.ID = int(BookingIDCount.ID + 1)

	// Id autoincrement payment
	PaymentIDCount := models.Payments{}
	configs.DB.Model(&models.Payments{}).Order("id desc").First(&PaymentIDCount)
	payment.ID = int(PaymentIDCount.ID + 1)

	// Calculate the total price
	ServicesPrice := models.Services{}
	configs.DB.Find(&ServicesPrice, "id = ?", booking.ServicesID)
	total := ServicesPrice.Price * float64(booking.Duration)
	booking.Total = total

	// Set the status to Awaiting Payment
	booking.Status = models.Awaiting_Payment

	// Set the created_at and updated_at
	booking.Created_At = time.Now()
	booking.Updated_At = time.Now()

	// Set the created_at and updated_at
	payment.Created_At = time.Now()
	payment.Updated_At = time.Now()

	// Set the payment id
	booking.PaymentsID = payment.ID

	if err := configs.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to create payment",
		})
	}

	if err := configs.DB.Create(&booking).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to create booking",
		})
	}

	// Update availibility on schedules table
	for i := booking.SchedulesID; i < booking.SchedulesID+booking.Duration; i++ {
		if err := configs.DB.Model(&models.Schedules{}).Where("id = ?", i).Update("Availbility", false).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Failed to update availibility",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create booking",
		"data":    booking,
	})

}

func UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	timeNow := time.Now()

	if err := configs.DB.Model(&models.Bookings{}).Where("id = ?", id).Update("status", models.Paid).Update("updated_at", timeNow).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update payment",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update payment",
	})

}

func UpdateBookingHousekeeper(c echo.Context) error {
	id := c.Param("id")
	timeNow := time.Now()

	if err := configs.DB.Model(&models.Bookings{}).Where("id = ?", id).Update("status", models.Cleaning).Update("updated_at", timeNow).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update booking",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update booking",
	})
}

func UpdateBookingHousekeeperDone(c echo.Context) error {
	id := c.Param("id")
	timeNow := time.Now()

	if err := configs.DB.Model(&models.Bookings{}).Where("id = ?", id).Update("status", models.Done).Update("updated_at", timeNow).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to update booking",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update booking",
	})
}
