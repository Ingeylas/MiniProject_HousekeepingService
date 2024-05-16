package controllers

import (
	"net/http"
	"rapikan/configs"
	mids "rapikan/middlewares"
	"rapikan/models"

	"time"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []models.Users

	if err := configs.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Users",
		"data":    users,
	})
}

func CreateUser(c echo.Context) error {
	user := models.Users{}

	c.Bind(&user)

	// check if username already exists
	var existingUser models.Users
	if err := configs.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Username already exists",
		})
	}

	// check if password has at least 8 characters
	if len(user.Password) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password must have at least 8 characters",
		})
	}

	user.Created_At = time.Now() // set time now
	user.Updated_At = time.Now() // set time now

	// // Check id and make autoincrement
	// idCount := int64(0)
	// configs.DB.Model(&models.Users{}).Where("ID = ?", user.ID).Count(&idCount)

	// if idCount > 0 {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "ID already exists",
	// 	})
	// }
	// user.ID = int(idCount + 1) // set id auto increment

	// if err := configs.DB.Create(&user).Error; err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "Failed to create user",
	// 	})
	// }

	// Check the last ID and add autoincrement
	UserIDCount := models.Users{}
	configs.DB.Model(&models.Users{}).Order("id desc").First(&UserIDCount)

	user.ID = int(UserIDCount.ID + 1)

	if err := configs.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Create User",
		"data":    user,
	})
}

func LoginUser(c echo.Context) error {
	user := models.Users{}

	c.Bind(&user)

	err := configs.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Not Found",
		})
	}

	token, err := mids.CreateToken(user.ID, user.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create token",
		})
	}

	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login User",
		"data":    userResponse,
	})

}
