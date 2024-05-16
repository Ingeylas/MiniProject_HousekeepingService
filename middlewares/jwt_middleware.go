package middlewares

import (
	"rapikan/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
