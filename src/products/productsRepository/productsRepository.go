package productsRepository

import (
	"E-Commerce/models/dto/productsDto"
	"E-Commerce/src/products"
	"database/sql"
	"github.com/lib/pq"
)

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) products.ProductsRepository {
	return productsRepository{db}
}

func (p productsRepository) CreateProducts(product productsDto.ProductsRepo) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	productsImage := pq.Array(product.ProductImage)
	productsQuery := `INSERT INTO
	  products (
		product_id,
		name,
		description,
		price,
		category_id,
		images
	  )
	VALUES
	  ($1, $2, $3, $4, $5, $6) RETURNING product_id`

	var productID string
	err = tx.QueryRow(productsQuery, product.ProductsID, product.ProductName, product.Description, product.Price, product.CategoryID, productsImage).Scan(&productID)
	if err != nil {
		return err
	}

	productsInventoryQuery := `INSERT INTO
	  inventory (product_id, stock_quantity)
	VALUES
	  ($1, $2)`

	_, err = tx.Exec(productsInventoryQuery, productID, product.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (p productsRepository) RetrieveALlProducts() ([]productsDto.ProductsResponse, error) {
	productsQuery := `SELECT
	  p.product_id,
	  p.name,
	  p.description,
	  p.price,
	  p.category_id,
	  p.images,
	  i.stock_quantity
	FROM
	  products p
	JOIN
	  inventory i
	ON
	  p.product_id = i.product_id`

	rows, err := p.db.Query(productsQuery)
	if err != nil {
		return nil, err
	}

	var products []productsDto.ProductsResponse
	for rows.Next() {
		var product productsDto.ProductsResponse
		var images pq.StringArray
		err := rows.Scan(&product.ProductsID, &product.ProductName, &product.Description, &product.Price, &product.CategoryID, &images, &product.Stock)
		if err != nil {
			return nil, err
		}
		product.ProductImage = images
		products = append(products, product)
	}

	return products, nil
}
