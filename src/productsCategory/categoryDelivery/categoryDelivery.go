package categoryDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsCategoryDto"
	"E-Commerce/pkg/middleware"
	"E-Commerce/src/productsCategory"
	"github.com/gin-gonic/gin"
)

type categoryDelivery struct {
	categoryUC productsCategory.CategoryUseCase
}

func NewCategoryDelivery(v1Group *gin.RouterGroup, categoryUC productsCategory.CategoryUseCase) {
	handler := categoryDelivery{
		categoryUC: categoryUC,
	}

	categoryGroup := v1Group.Group("/categories")
	{
		categoryGroup.POST("", middleware.JWTAuth("admin"), handler.CreateCategory)
	}
}

func (cat categoryDelivery) CreateCategory(ctx *gin.Context) {
	var req productsCategoryDto.ProductsCategoryReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	categoryData, err := cat.categoryUC.CreateCategory(req.CategoryName)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, categoryData, nil, "success create category", constants.ServiceCodeCategory, constants.SuccessCode)
}
