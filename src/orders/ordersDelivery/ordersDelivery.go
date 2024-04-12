package ordersDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/ordersDto"
	"E-Commerce/src/orders"
	"github.com/gin-gonic/gin"
)

type ordersDelivery struct {
	ordersUC orders.OrdersUsecase
}

func NewOrdersDelivery(v1Group *gin.RouterGroup, ordersUC orders.OrdersUsecase) {
	handler := ordersDelivery{
		ordersUC: ordersUC,
	}

	ordersGroup := v1Group.Group("/orders")
	{
		ordersGroup.POST("", handler.CreateOrder)
	}

}

func (o ordersDelivery) CreateOrder(ctx *gin.Context) {
	var req ordersDto.OrderReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeOrder, constants.GeneralErrCode)
		return
	}

	order, err := o.ordersUC.CreateOrder(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeOrder, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, order, nil, "Success Create Order", constants.ServiceCodeOrder, constants.SuccessCode)
}
