package categoryUseCase

import (
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsCategoryDto"
	"E-Commerce/src/productsCategory"
	"errors"
	"github.com/google/uuid"
	"math"
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

func (c categoryUC) RetrieveAllCategory(page, pageSize int) ([]productsCategoryDto.ProductsCategoryDto, json.Pagination, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	categoryData, err := c.categoryRepo.RetrieveAllCategory(page, pageSize)
	if err != nil {
		return nil, json.Pagination{}, err
	}

	totalCategoryData, err := c.categoryRepo.CountAllCategory()
	if err != nil {
		return nil, json.Pagination{}, err
	}

	totalPages := int(math.Ceil(float64(totalCategoryData) / float64(pageSize)))
	if page > totalPages {
		return nil, json.Pagination{}, errors.New("01")
	}

	if totalPages == 0 && totalCategoryData > 0 {
		totalPages = 1
	}

	pagination := json.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: totalCategoryData,
	}

	return categoryData, pagination, nil
}

func (c categoryUC) RetrieveCategoryById(categoryId string) (productsCategoryDto.ProductsCategoryDto, error) {
	categoryData, err := c.categoryRepo.RetrieveCategoryById(categoryId)
	if err != nil {
		if err.Error() == "01" {
			return productsCategoryDto.ProductsCategoryDto{}, errors.New("01")
		}
		return productsCategoryDto.ProductsCategoryDto{}, err
	}

	return categoryData, nil
}

func (c categoryUC) UpdateCategory(categoryId, categoryName string) error {
	categoryData, err := c.categoryRepo.RetrieveCategoryById(categoryId)
	if err != nil {
		if err.Error() == "01" {
			return errors.New("01")
		}
		return err
	}

	catExists, err := c.categoryRepo.CategoryExist(categoryName)
	if err != nil {
		return err
	}

	if catExists {
		return errors.New("02")
	}

	if categoryName != "" {
		categoryData.CategoryName = categoryName
	}

	err = c.categoryRepo.UpdateCategory(categoryId, categoryData.CategoryName)
	if err != nil {
		return err
	}

	return nil
}

func (c categoryUC) DeleteCategory(categoryId string) error {
	categoryData, err := c.categoryRepo.RetrieveCategoryById(categoryId)
	if err != nil {
		if err.Error() == "01" {
			return errors.New("01")
		}
		return err
	}

	catIDStr := categoryData.CategoryId.String()
	err = c.categoryRepo.DeleteCategory(catIDStr)
	if err != nil {
		return err
	}

	return nil
}
