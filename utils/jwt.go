package utils

import (
	"fmt"
	"go-practice/HOTEL/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("LEARNING GO ")

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Name,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(accessToken)

	return accessToken, nil

}
