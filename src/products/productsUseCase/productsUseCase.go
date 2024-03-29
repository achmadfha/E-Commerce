package productsUseCase

import (
	"E-Commerce/models/dto/productsDto"
	"E-Commerce/pkg/utils"
	"E-Commerce/src/products"
	"errors"
	"github.com/google/uuid"
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

func (prod productsUC) CreateProducts(product productsDto.ProductsRequest) (productsDto.ProductsResponse, error) {
	productsID, err := uuid.NewRandom()
	if err != nil {
		return productsDto.ProductsResponse{}, err
	}

	//Todo
	// check if categoryID exists in the database
	// if not exists return error 'category not found
	// if exists continue to the next process
	// check if there is a product with the same name in the database
	// if exists return error 'product name already exists'
	// if not exists continue to the next process

	productRepo := productsDto.ProductsRepo{
		ProductsID:   productsID,
		ProductName:  product.ProductName,
		ProductImage: product.ProductImage,
		Description:  product.Description,
		Price:        product.Price,
		CategoryID:   product.CategoryID,
		Stock:        product.Stock,
	}

	err = prod.productsRepo.CreateProducts(productRepo)
	if err != nil {
		return productsDto.ProductsResponse{}, err
	}

	productResp := productsDto.ProductsResponse{
		ProductsID:   productsID,
		ProductName:  productRepo.ProductName,
		ProductImage: productRepo.ProductImage,
		Description:  productRepo.Description,
		Price:        productRepo.Price,
		CategoryID:   productRepo.CategoryID,
		Stock:        productRepo.Stock,
	}

	return productResp, nil
}
