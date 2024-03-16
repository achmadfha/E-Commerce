package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func ExtractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return ""
	}

	return authHeaderParts[1]
}

func ParseTokenAndExtractClaims(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("SECRET_TOKEN")
	key := []byte(secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
