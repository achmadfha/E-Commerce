package productsCategoryDto

import "github.com/google/uuid"

type (
	ProductsCategoryDto struct {
		CategoryId   uuid.UUID `json:"category_id"`
		CategoryName string    `json:"name"`
	}

	ProductsCategoryReq struct {
		CategoryName string `json:"name"`
	}
)
