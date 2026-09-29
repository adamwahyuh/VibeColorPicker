package utils

import (
	"time"

	"github.com/adamwahyuh/coloroutfit/src/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(UserId uint16) (string, error) {
	config.LoadEnv()

	jwtSecret := []byte(config.GetEnv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"user_id": UserId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtSecret)
}
