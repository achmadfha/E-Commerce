package middleware

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func JWTAuth(roles ...string) gin.HandlerFunc {
	secret := os.Getenv("SECRET_TOKEN")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseUnauthorized(c, "Unauthorized. [Invalid Token]", constants.ServiceCodeJWT, constants.Unauthorized)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			json.NewResponseUnauthorized(c, "Unauthorized. [Invalid Token]", constants.ServiceCodeJWT, constants.Unauthorized)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			json.NewResponseUnauthorized(c, "Unauthorized. [Invalid Token]", constants.ServiceCodeJWT, constants.Unauthorized)
			c.Abort()
			return
		}

		expirationTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expirationTime {
			json.NewResponseUnauthorized(c, "Unauthorized. [Token Expired]", constants.ServiceCodeJWT, constants.Unauthorized)
			c.Abort()
			return
		}

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims["role"] {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			json.NewResponseForbidden(c, "Unauthorized. [Forbidden]", constants.ServiceCodeJWT, constants.Unauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
