package router

import (
	"E-Commerce/src/authentication/authenticationDelivery"
	"E-Commerce/src/authentication/authenticationRepository"
	"E-Commerce/src/authentication/authenticationUseCase"
	"E-Commerce/src/orders/ordersDelivery"
	"E-Commerce/src/orders/ordersRepository"
	"E-Commerce/src/orders/ordersUseCase"
	"E-Commerce/src/products/productsDelivery"
	"E-Commerce/src/products/productsRepository"
	"E-Commerce/src/products/productsUseCase"
	"E-Commerce/src/productsCategory/categoryDelivery"
	"E-Commerce/src/productsCategory/categoryRepository"
	"E-Commerce/src/productsCategory/categoryUseCase"
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
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	productsRepo := productsRepository.NewProductsRepository(db)
	orderRepo := ordersRepository.NewOrdersRepository(db)

	// usecase
	authenticationUC := authenticationUseCase.NewAuthenticationUseCase(authenticationRepo)
	usersUC := usersUseCase.NewUserUseCase(usersRepo)
	categoryUC := categoryUseCase.NewCategoryUseCase(categoryRepo)
	productsUC := productsUseCase.NewProductsUseCase(productsRepo, categoryRepo)
	ordersUC := ordersUseCase.NewOrdersUseCase(orderRepo)

	// delivery
	authenticationDelivery.NewAuthenticationDelivery(v1Group, authenticationUC)
	usersDelivery.NewUserDelivery(v1Group, usersUC)
	categoryDelivery.NewCategoryDelivery(v1Group, categoryUC)
	productsDelivery.NewProductsDelivery(v1Group, productsUC)
	ordersDelivery.NewOrdersDelivery(v1Group, ordersUC)
}
