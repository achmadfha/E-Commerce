package products

import "io"

type ProductsRepository interface {
}

type ProductsUseCase interface {
	UploadProductsImages(fileContent io.Reader) (string, error)
}
