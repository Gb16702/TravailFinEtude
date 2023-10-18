package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(email string, Id uint) (string, error) {
	var key = []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["authorized"] = true
	claims["user"] = email
	claims["id"] = Id

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
