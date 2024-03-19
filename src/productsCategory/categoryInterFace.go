package productsCategory

import "E-Commerce/models/dto/productsCategoryDto"

type CategoryRepository interface {
	CreateCategory(prod productsCategoryDto.ProductsCategoryDto) error
	CategoryExist(categoryName string) (bool, error)
}

type CategoryUseCase interface {
	CreateCategory(categoryName string) (cat productsCategoryDto.ProductsCategoryDto, err error)
}
