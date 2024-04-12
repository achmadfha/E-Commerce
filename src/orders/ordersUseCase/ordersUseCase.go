package ordersUseCase

import (
	"E-Commerce/models/dto/ordersDto"
	"E-Commerce/src/orders"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type ordersUseCase struct {
	ordersRepo orders.OrdersRepository
}

func NewOrdersUseCase(ordersRepo orders.OrdersRepository) orders.OrdersUsecase {
	return &ordersUseCase{ordersRepo}
}

func (o ordersUseCase) CreateOrder(order ordersDto.OrderReq) (ordersDto.OrderResp, error) {
	orderID, err := uuid.NewUUID()
	if err != nil {
		return ordersDto.OrderResp{}, err
	}

	shippingID, err := uuid.NewUUID()
	if err != nil {
		return ordersDto.OrderResp{}, err
	}

	trackingNumber := fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(100))
	shippingCarriers := []string{"UPS", "FedEx", "DHL", "USPS"}
	shippingCarrier := shippingCarriers[rand.Intn(len(shippingCarriers))]
	shippingCost := 10000
	totalAmount := 0.0
	for _, item := range order.Products {
		totalAmount += float64(item.Quantity) * item.Price
	}
	totalAmount += float64(shippingCost)
	orderRepo := ordersDto.OrderRepo{
		UsersID:         order.UsersID,
		OrderID:         orderID,
		TotalAmount:     totalAmount,
		Status:          "pending",
		OrderItems:      order.Products,
		ShippingID:      shippingID,
		TrackingNumber:  trackingNumber,
		ShippingCarrier: shippingCarrier,
		ShippingCost:    float64(shippingCost),
		ShippingAddress: order.Shipping,
	}

	err = o.ordersRepo.CreateOrder(orderRepo)
	if err != nil {
		return ordersDto.OrderResp{}, err
	}

	orderResp := ordersDto.OrderResp{
		OrderID:     orderID,
		ShipmentsID: shippingID,
		TotalAmount: totalAmount,
		Status:      "pending",
	}

	return orderResp, nil
}
