package utils

import (
	"fmt"
	"go-practice/HOTEL/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"phone":    user.Phone,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(accessToken)

	return accessToken, nil

}

func ParseToken(accessToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	return token, err
}
