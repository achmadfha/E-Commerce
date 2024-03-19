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
	query := `INSERT INTO
	  categories (category_id, name)
	VALUES
	  ($1, $2)`

	_, err := p.db.Exec(query, prod.CategoryId, prod.CategoryName)
	if err != nil {
		return err
	}
	return nil
}

func (p productsRepository) CategoryExist(categoryName string) (bool, error) {
	query := `SELECT
	  EXISTS(
		SELECT
		  1
		FROM
		  categories
		WHERE
		  name ILIKE $1
	  )`

	var exists bool
	err := p.db.QueryRow(query, categoryName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
