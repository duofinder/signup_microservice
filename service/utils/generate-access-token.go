package serviceutils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"sub":     userId,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * 1).Unix(),
		"iss":     "duofinder.app",
		"isAdmin": false,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := jwt.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
