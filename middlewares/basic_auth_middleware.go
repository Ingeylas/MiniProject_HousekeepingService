package middlewares

import (
	"rapikan/configs"
	"rapikan/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user models.Users
	err := configs.DB.Where("username = ? AND password = ?", username, password).First(&user).Error

	if err != nil {
		return false, err
	}
	return true, nil

}
