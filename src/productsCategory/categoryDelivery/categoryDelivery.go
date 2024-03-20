package categoryDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsCategoryDto"
	"E-Commerce/pkg/middleware"
	"E-Commerce/src/productsCategory"
	"github.com/gin-gonic/gin"
	"strconv"
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
		categoryGroup.GET("", middleware.JWTAuth("admin", "users"), handler.RetrieveAllCategory)
		categoryGroup.GET("/:categoryId", middleware.JWTAuth("admin", "users"), handler.RetrieveCategoryById)
		categoryGroup.PUT("/:categoryId", middleware.JWTAuth("admin"), handler.UpdateCategoryName)
		categoryGroup.DELETE("/:categoryId", middleware.JWTAuth("admin"), handler.DeleteCategoryById)
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
		if err.Error() == "01" {
			json.NewResponseError(ctx, "category already exists", constants.ServiceCodeCategory, constants.GeneralErrCode)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, categoryData, nil, "success create category", constants.ServiceCodeCategory, constants.SuccessCode)
}

func (cat categoryDelivery) RetrieveAllCategory(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("size"))

	categoryData, pagination, err := cat.categoryUC.RetrieveAllCategory(page, pageSize)
	if err != nil {
		if err.Error() == "01" {
			errorMessage := "page " + strconv.Itoa(page) + " doesn't exist"
			json.NewResponseForbidden(ctx, errorMessage, constants.ServiceCodeCategory, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, categoryData, pagination, "success retrieve all category", constants.ServiceCodeCategory, constants.SuccessCode)
}

func (cat categoryDelivery) RetrieveCategoryById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")

	categoryData, err := cat.categoryUC.RetrieveCategoryById(categoryId)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "category doesn't exist", constants.ServiceCodeCategory, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, categoryData, nil, "success retrieve category by id", constants.ServiceCodeCategory, constants.SuccessCode)
}

func (cat categoryDelivery) UpdateCategoryName(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	var req productsCategoryDto.ProductsCategoryReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	err := cat.categoryUC.UpdateCategory(categoryId, req.CategoryName)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "category doesn't exist", constants.ServiceCodeCategory, constants.Forbidden)
			return
		}
		if err.Error() == "02" {
			json.NewResponseForbidden(ctx, "category already exist exist", constants.ServiceCodeCategory, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "success update category name", constants.ServiceCodeCategory, constants.SuccessCode)
}

func (cat categoryDelivery) DeleteCategoryById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")

	err := cat.categoryUC.DeleteCategory(categoryId)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "category doesn't exist", constants.ServiceCodeCategory, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeCategory, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "success delete category by id", constants.ServiceCodeCategory, constants.SuccessCode)
}
