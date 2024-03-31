package products

import (
	"E-Commerce/models/dto/productsDto"
	"io"
)

type ProductsRepository interface {
	CreateProducts(product productsDto.ProductsRepo) error
	RetrieveALlProducts() ([]productsDto.ProductsResponse, error)
	RetrieveProductByID(productID string) (productsDto.ProductsResponseDetail, error)
	UpdateProducts(product productsDto.ProductsRepo) error
	DeleteProducts(productID string) error
}

type ProductsUseCase interface {
	UploadProductsImages(fileContent io.Reader) (string, error)
	CreateProducts(product productsDto.ProductsRequest) (productsDto.ProductsResponse, error)
	RetrieveAllProducts() ([]productsDto.ProductsResponse, error)
	RetrieveProductByID(productID string) (productsDto.ProductsResponseDetail, error)
	UpdateProducts(product productsDto.ProductsRepo) error
	DeleteProducts(productID string) error
}
