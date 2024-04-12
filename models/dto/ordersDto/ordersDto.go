package ordersDto

import "github.com/google/uuid"

type (
	OrderReq struct {
		UsersID  uuid.UUID       `json:"user_id"`
		Products []OrderItems    `json:"products"`
		Shipping ShippingAddress `json:"shipping_address"`
	}

	OrderItems struct {
		ProductsID uuid.UUID `json:"products_id"`
		Quantity   int       `json:"quantity"`
		Price      float64   `json:"price"`
	}

	ShippingAddress struct {
		Address    string `json:"address"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
	}

	OrderRepo struct {
		UsersID         uuid.UUID       `json:"user_id"`
		OrderID         uuid.UUID       `json:"order_id"`
		TotalAmount     float64         `json:"total_amount"`
		Status          string          `json:"status"`
		OrderItems      []OrderItems    `json:"order_items"`
		ShippingID      uuid.UUID       `json:"shipping_id"`
		TrackingNumber  string          `json:"tracking_number"`
		ShippingCarrier string          `json:"shipping_carrier"`
		ShippingCost    float64         `json:"shipping_cost"`
		ShippingAddress ShippingAddress `json:"shipping_address"`
	}

	OrderResp struct {
		OrderID     uuid.UUID `json:"order_id"`
		ShipmentsID uuid.UUID `json:"shipment_id"`
		TotalAmount float64   `json:"total_amount"`
		Status      string    `json:"status"`
	}
)
