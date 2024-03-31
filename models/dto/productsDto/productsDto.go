package productsDto

import "github.com/google/uuid"

type (
	ProductsRequest struct {
		ProductName  string    `json:"product_name"`
		ProductImage []string  `json:"product_image"`
		Description  string    `json:"description"`
		Price        int       `json:"price"`
		CategoryID   uuid.UUID `json:"category_id"`
		Stock        int       `json:"stock"`
	}

	ProductsRepo struct {
		ProductsID   uuid.UUID `json:"products_id"`
		ProductName  string    `json:"name"`
		ProductImage []string  `json:"image"`
		Description  string    `json:"description"`
		Price        int       `json:"price"`
		CategoryID   uuid.UUID `json:"category_id"`
		Stock        int       `json:"stock"`
	}

	ProductsResponse struct {
		ProductsID   uuid.UUID `json:"products_id"`
		ProductName  string    `json:"product_name"`
		ProductImage []string  `json:"product_image"`
		Description  string    `json:"description"`
		Price        float64   `json:"price"`
		CategoryID   uuid.UUID `json:"category_id"`
		Stock        int       `json:"stock"`
	}
)
