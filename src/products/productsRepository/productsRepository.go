package productsRepository

import (
	"E-Commerce/models/dto/productsCategoryDto"
	"E-Commerce/src/productsCategory"
	"database/sql"
)

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) productsCategory.CategoryRepository {
	return productsRepository{db}
}

func (p productsRepository) CreateCategory(prod productsCategoryDto.ProductsCategoryDto) error {
	query := `INSERT INTO categories (category_id, name) VALUES ($1, $2)`

	_, err := p.db.Exec(query, prod.CategoryId, prod.CategoryName)
	if err != nil {
		return err
	}
	return nil
}
