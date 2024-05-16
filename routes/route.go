package routes

import (
	"rapikan/controllers"
	mids "rapikan/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"

	"rapikan/constants"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsers)       // Get All Users
	e.POST("/user", controllers.CreateUser)     // Register User
	e.POST("/loginUser", controllers.LoginUser) // Login User

	e.GET("/services", controllers.GetServices)   // Get All Services
	e.POST("/service", controllers.CreateService) // Add Service

	e.POST("/housekeeper", controllers.CreateHousekeeper)      // Add Housekeeper
	e.GET("/housekeepers", controllers.GetHousekeepers)        // Get All Housekeepers
	e.GET("/housekeepers/:id", controllers.GetHousekeeperByID) // Get Housekeeper by ServiceID

	e.POST("/schedule/:id", controllers.CreateSchedule)                         // Auto Generate Schedule on 1 day
	e.GET("/schedules", controllers.GetSchedules)                               // Get All Schedules
	e.GET("/schedules/:id", controllers.GetScheduleByID)                        // Get Schedule by ID
	e.GET("/bookings/housekeepers/:id", controllers.GetBookingsByHousekeeperID) // Get Bookings by HousekeeperID

	e.GET("/bookings", controllers.GetBookings)                                      // Get All Bookings
	e.POST("/booking", controllers.CreateBooking)                                    // Create Booking
	e.PUT("/booking/:id", controllers.UpdatePayment)                                 // Update Booking Status based on payment status
	e.PUT("/booking/housekeeper/:id", controllers.UpdateBookingHousekeeper)          // Update Booking Status based on housekeeper status
	e.PUT("/booking/housekeeper/:id/done", controllers.UpdateBookingHousekeeperDone) // Update Booking Status to Done

	e.GET("/payments", controllers.GetPayments)   // Get All Payments
	e.POST("/payment", controllers.CreatePayment) // Create Payment

	eAuthBasic := e.Group("/login")
	eAuthBasic.Use(mid.BasicAuth(mids.BasicAuthDB))
	eAuthBasic.GET("/user", controllers.GetUsers)

	eAuthJWT := e.Group("/jwt")
	eAuthJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eAuthJWT.GET("/user", controllers.GetUsers)

	return e
}
