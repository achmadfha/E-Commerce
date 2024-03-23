package productsUseCase

import (
	"E-Commerce/pkg/utils"
	"E-Commerce/src/products"
	"errors"
	"io"
	"net/http"
)

type productsUC struct {
	productsRepo products.ProductsRepository
}

func NewProductsUseCase(productsRepo products.ProductsRepository) products.ProductsUseCase {
	return &productsUC{productsRepo}
}

func (prod productsUC) UploadProductsImages(fileContent io.Reader) (string, error) {
	if fileContent == nil {
		return "", errors.New("file content is nil")
	}

	fileHeader := make([]byte, 512)
	_, err := fileContent.Read(fileHeader)
	if err != nil {
		return "", err
	}

	if _, err := fileContent.(io.Seeker).Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	fileType := http.DetectContentType(fileHeader)
	if fileType != "image/jpeg" && fileType != "image/png" {
		return "", errors.New("file is not a valid image")
	}

	urls, err := utils.UploadToS3(fileContent)
	if err != nil {
		return "", err
	}

	return urls, nil
}
