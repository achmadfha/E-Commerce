package categoryRepository

import (
	"E-Commerce/models/dto/productsCategoryDto"
	"E-Commerce/src/productsCategory"
	"database/sql"
	"errors"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) productsCategory.CategoryRepository {
	return categoryRepository{db}
}

func (c categoryRepository) CreateCategory(prod productsCategoryDto.ProductsCategoryDto) error {
	query := `INSERT INTO
	  categories (category_id, name)
	VALUES
	  ($1, $2)`

	_, err := c.db.Exec(query, prod.CategoryId, prod.CategoryName)
	if err != nil {
		return err
	}
	return nil
}

func (c categoryRepository) CategoryExist(categoryName string) (bool, error) {
	query := `SELECT
	  EXISTS(
		SELECT
		  1
		FROM
		  categories
		WHERE
		  name ILIKE $1
		AND is_deleted = false
	  )`

	var exists bool
	err := c.db.QueryRow(query, categoryName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (c categoryRepository) RetrieveAllCategory(page, pageSize int) ([]productsCategoryDto.ProductsCategoryDto, error) {
	query := `SELECT
	  category_id,
	  name
	FROM
	  categories
	WHERE
	  is_deleted = false
	LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []productsCategoryDto.ProductsCategoryDto
	for rows.Next() {
		var category productsCategoryDto.ProductsCategoryDto
		err := rows.Scan(&category.CategoryId, &category.CategoryName)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c categoryRepository) CountAllCategory() (int, error) {
	query := `SELECT
	  COUNT(*)
	FROM
	  categories
	WHERE
	  is_deleted = false`

	var count int
	err := c.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c categoryRepository) RetrieveCategoryById(categoryId string) (productsCategoryDto.ProductsCategoryDto, error) {
	query := `SELECT
	  category_id,
	  name
	FROM
	  categories
	WHERE
	  category_id = $1
	AND 
	  is_deleted = false`

	var category productsCategoryDto.ProductsCategoryDto
	err := c.db.QueryRow(query, categoryId).Scan(&category.CategoryId, &category.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return productsCategoryDto.ProductsCategoryDto{}, errors.New("01")
		}
		return productsCategoryDto.ProductsCategoryDto{}, err
	}

	return category, nil
}

func (c categoryRepository) UpdateCategory(categoryId, categoryName string) error {
	query := `UPDATE
	  categories
	SET
	  name = $1
	WHERE
	  category_id = $2
	AND
	  is_deleted = false`

	_, err := c.db.Exec(query, categoryName, categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (c categoryRepository) DeleteCategory(categoryId string) error {
	query := `UPDATE
	  categories
	SET
	  is_deleted = true
	WHERE
	  category_id = $1`

	_, err := c.db.Exec(query, categoryId)
	if err != nil {
		return err
	}

	return nil
}
