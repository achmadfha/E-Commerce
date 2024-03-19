package categoryUseCase

import (
	"E-Commerce/models/dto/productsCategoryDto"
	"E-Commerce/src/productsCategory"
	"errors"
	"github.com/google/uuid"
)

type categoryUC struct {
	categoryRepo productsCategory.CategoryRepository
}

func NewCategoryUseCase(categoryRepo productsCategory.CategoryRepository) productsCategory.CategoryUseCase {
	return &categoryUC{categoryRepo}
}

func (c categoryUC) CreateCategory(categoryName string) (cat productsCategoryDto.ProductsCategoryDto, err error) {
	catID, err := uuid.NewRandom()

	catExists, err := c.categoryRepo.CategoryExist(categoryName)
	if err != nil {
		return productsCategoryDto.ProductsCategoryDto{}, err
	}

	if catExists {
		return productsCategoryDto.ProductsCategoryDto{}, errors.New("01")
	}

	prodData := productsCategoryDto.ProductsCategoryDto{
		CategoryId:   catID,
		CategoryName: categoryName,
	}

	err = c.categoryRepo.CreateCategory(prodData)
	if err != nil {
		return productsCategoryDto.ProductsCategoryDto{}, err
	}

	return prodData, nil
}
