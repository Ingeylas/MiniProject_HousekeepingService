package controllers

import (
	"fmt"
	"net/http"
	"rapikan/configs"
	"rapikan/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// func CreateSchedule(c echo.Context) error {
// 	schedule := models.Schedules{}

// 	c.Bind(&schedule)

// 	// check if schedule already exists
// 	var existingSchedule models.Schedules
// 	if err := configs.DB.Where("service_id = ?", schedule.ID).First(&existingSchedule).Error; err == nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "Schedule already exists",
// 		})
// 	}

// 	schedule.Date = time.Now() // set time now
// 	// schedule.Starting_Hour = time.Date(0, 1, 1, 17, 0, 0, 0, time.UTC)

// 	// Check the last ID and add autoincrement
// 	ScheduleIDCount := models.Schedules{}
// 	configs.DB.Model(&models.Schedules{}).Order("id desc").First(&ScheduleIDCount)

// 	schedule.ID = int(ScheduleIDCount.ID + 1)

// 	if err := configs.DB.Create(&schedule).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "Failed to create schedule",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Success create schedule",
// 		"data":    schedule,
// 	})
// }

func GetSchedules(c echo.Context) error {
	var schedules []models.Schedules

	if err := configs.DB.Preload("Housekeepers.Services").Find(&schedules).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Schedules",
		"data":    schedules,
	})
}

func CreateSchedule(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i := 0; i < 7; i++ {
		schedule := models.Schedules{}

		//Id autoincrement
		scheduleIDCount := models.Schedules{}
		configs.DB.Model(&models.Schedules{}).Order("id desc").First(&scheduleIDCount)
		schedule.ID = int(scheduleIDCount.ID + 1)

		//Set Time Now
		schedule.Date = time.Now() // set time now

		//Set Working Hour
		schedule.Working_Hour = fmt.Sprintf("%d:00", 10+i)

		//Set Availbility
		schedule.Availbility = true

		//Set HousekeepersID
		schedule.HousekeepersID = id

		if err := configs.DB.Create(&schedule).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Failed to create schedule",
			})
		}

		// return c.JSON(http.StatusOK, map[string]interface{}{
		// 	"message": "Success create schedule",
		// 	"data":    schedule,
		// })

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create schedule",
	})

}

func GetScheduleByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var schedules []models.Schedules

	if err := configs.DB.Preload("Housekeepers.Services").Where("housekeepers_id = ?", id).Find(&schedules).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Schedules",
		"data":    schedules,
	})
}
