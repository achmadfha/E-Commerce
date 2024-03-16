package router

import (
	"E-Commerce/src/authentication/authenticationDelivery"
	"E-Commerce/src/authentication/authenticationRepository"
	"E-Commerce/src/authentication/authenticationUseCase"
	"E-Commerce/src/users/usersDelivery"
	"E-Commerce/src/users/usersRepository"
	"E-Commerce/src/users/usersUseCase"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authenticationRepo := authenticationRepository.NewAuthenticationRepository(db)
	usersRepo := usersRepository.NewUserRepository(db)

	// usecase
	authenticationUC := authenticationUseCase.NewAuthenticationUseCase(authenticationRepo)
	usersUC := usersUseCase.NewUserUseCase(usersRepo)

	// delivery
	authenticationDelivery.NewAuthenticationDelivery(v1Group, authenticationUC)
	usersDelivery.NewUserDelivery(v1Group, usersUC)
}
