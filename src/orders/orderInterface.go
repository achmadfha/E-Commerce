package orders

import "E-Commerce/models/dto/ordersDto"

type OrdersRepository interface {
	CreateOrder(order ordersDto.OrderRepo) error
}

type OrdersUsecase interface {
	CreateOrder(order ordersDto.OrderReq) (ordersDto.OrderResp, error)
}
