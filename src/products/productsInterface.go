package products

import (
	"E-Commerce/models/dto/productsDto"
	"io"
)

type ProductsRepository interface {
	CreateProducts(product productsDto.ProductsRepo) error
}

type ProductsUseCase interface {
	UploadProductsImages(fileContent io.Reader) (string, error)
	CreateProducts(product productsDto.ProductsRequest) (productsDto.ProductsResponse, error)
}
