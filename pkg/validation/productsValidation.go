package validation

import (
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsDto"
	"github.com/google/uuid"
)

func ValidationProducts(req productsDto.ProductsRequest) []json.ValidationField {
	var validationErrors []json.ValidationField

	if req.ProductName == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "product_name",
			Message:   "Product name cannot be empty",
		})
	}

	if len(req.ProductImage) == 0 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "product_image",
			Message:   "Product image cannot be empty",
		})
	}

	if req.Description == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "description",
			Message:   "Description cannot be empty",
		})
	}

	if req.Price == 0 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "price",
			Message:   "Price cannot be empty",
		})
	}

	if req.CategoryID == uuid.Nil {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "category_id",
			Message:   "Category ID cannot be empty",
		})
	}

	if req.Stock == 0 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "stock",
			Message:   "Stock cannot be empty",
		})
	}

	return validationErrors
}
