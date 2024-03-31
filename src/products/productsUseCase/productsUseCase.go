package productsUseCase

import (
	"E-Commerce/models/dto/productsDto"
	"E-Commerce/pkg/utils"
	"E-Commerce/src/products"
	"E-Commerce/src/productsCategory"
	"errors"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type productsUC struct {
	productsRepo products.ProductsRepository
	categoryRepo productsCategory.CategoryRepository
}

func NewProductsUseCase(productsRepo products.ProductsRepository, categoryRepo productsCategory.CategoryRepository) products.ProductsUseCase {
	return &productsUC{productsRepo, categoryRepo}
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

	// Check the file size
	const maxFileSize = 2 << 20 // 2MB
	limitedReader := io.LimitedReader{R: fileContent, N: maxFileSize + 1}
	_, err = io.CopyN(io.Discard, &limitedReader, maxFileSize+1)
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}

	if limitedReader.N <= 0 {
		return "", errors.New("file size is more than 2MB")
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

	catID := product.CategoryID.String()
	categoryData, err := prod.categoryRepo.RetrieveCategoryById(catID)
	if err != nil {
		if err.Error() == "01" {
			return productsDto.ProductsResponse{}, errors.New("01")
		}
		return productsDto.ProductsResponse{}, err
	}

	productRepo := productsDto.ProductsRepo{
		ProductsID:   productsID,
		ProductName:  product.ProductName,
		ProductImage: product.ProductImage,
		Description:  product.Description,
		Price:        product.Price,
		CategoryID:   categoryData.CategoryId,
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
		Price:        float64(productRepo.Price),
		CategoryID:   categoryData.CategoryId,
		Stock:        productRepo.Stock,
	}

	return productResp, nil
}

func (prod productsUC) RetrieveAllProducts() ([]productsDto.ProductsResponse, error) {
	productsData, err := prod.productsRepo.RetrieveALlProducts()
	if err != nil {
		return nil, err
	}

	var productsResponse []productsDto.ProductsResponse
	for _, product := range productsData {
		productsResponse = append(productsResponse, productsDto.ProductsResponse{
			ProductsID:   product.ProductsID,
			ProductName:  product.ProductName,
			ProductImage: product.ProductImage,
			Description:  product.Description,
			Price:        product.Price,
			CategoryID:   product.CategoryID,
			Stock:        product.Stock,
		})
	}

	return productsResponse, nil
}
