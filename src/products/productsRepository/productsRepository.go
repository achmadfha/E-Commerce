package productsRepository

import (
	"E-Commerce/src/products"
	"database/sql"
)

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) products.ProductsRepository {
	return productsRepository{db}
}
