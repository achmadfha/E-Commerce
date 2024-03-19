package productsCategory

import (
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsCategoryDto"
)

type CategoryRepository interface {
	CreateCategory(prod productsCategoryDto.ProductsCategoryDto) error
	CategoryExist(categoryName string) (bool, error)
	RetrieveAllCategory(page, pageSize int) ([]productsCategoryDto.ProductsCategoryDto, error)
	CountAllCategory() (int, error)
	RetrieveCategoryById(categoryId string) (productsCategoryDto.ProductsCategoryDto, error)
}

type CategoryUseCase interface {
	CreateCategory(categoryName string) (cat productsCategoryDto.ProductsCategoryDto, err error)
	RetrieveAllCategory(page, pageSize int) ([]productsCategoryDto.ProductsCategoryDto, json.Pagination, error)
	RetrieveCategoryById(categoryId string) (productsCategoryDto.ProductsCategoryDto, error)
}
