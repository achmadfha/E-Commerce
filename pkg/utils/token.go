package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

func GenerateToken(id uuid.UUID, role string) (tokenString string, err error) {
	expired := os.Getenv("TOKEN_EXPIRED")
	secret := os.Getenv("SECRET_TOKEN")
	exp, err := strconv.Atoi(expired)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiredTime := time.Now().Add(time.Duration(exp) * time.Hour)

	claims := jwt.MapClaims{
		"clientId": id,
		"exp":      expiredTime.Unix(),
		"role":     role,
		"iat":      now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
