package router

import (
	"E-Commerce/src/authentication/authenticationDelivery"
	"E-Commerce/src/authentication/authenticationRepository"
	"E-Commerce/src/authentication/authenticationUseCase"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authenticationRepo := authenticationRepository.NewAuthenticationRepository(db)

	// usecase
	authenticationUC := authenticationUseCase.NewAuthenticationUseCase(authenticationRepo)

	// delivery
	authenticationDelivery.NewAuthenticationDelivery(v1Group, authenticationUC)
}
