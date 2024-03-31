package productsDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsDto"
	"E-Commerce/pkg/validation"
	"E-Commerce/src/products"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
)

type productsDelivery struct {
	productsUC products.ProductsUseCase
}

func NewProductsDelivery(v1Group *gin.RouterGroup, productsUC products.ProductsUseCase) {
	handler := productsDelivery{
		productsUC: productsUC,
	}

	productsGroup := v1Group.Group("/products")
	{
		productsGroup.POST("/upload-images", handler.UploadProductsImages)
		productsGroup.POST("", handler.CreateProducts)
		productsGroup.GET("", handler.RetrieveAllProducts)
		productsGroup.GET("/:id", handler.RetrieveProductsByID)
		productsGroup.PUT("/:id", handler.UpdateProducts)
		productsGroup.DELETE("/:id", handler.DeleteProducts)
	}

}

func (prod productsDelivery) UploadProductsImages(ctx *gin.Context) {
	file, err := ctx.FormFile("fileName")
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {
			json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
			return
		}
	}(fileContent)

	url, err := prod.productsUC.UploadProductsImages(fileContent)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	data := interface{}(map[string]string{"url": url})
	json.NewResponseSuccess(ctx, data, nil, "Success Upload Image", constants.ServiceCodeProduct, constants.SuccessCode)
}

func (prod productsDelivery) CreateProducts(ctx *gin.Context) {
	var req productsDto.ProductsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	validationErrors := validation.ValidationProducts(req)
	if len(validationErrors) > 0 {
		json.NewResponseBadRequest(ctx, validationErrors, constants.BadReqMsg, constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	product, err := prod.productsUC.CreateProducts(req)
	if err != nil {
		if err.Error() == "01" {
			msg := fmt.Sprintf("category with id %s not found", req.CategoryID.String())
			json.NewResponseError(ctx, msg, constants.ServiceCodeProduct, constants.NotFoundCode)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, product, nil, "Success Create Product", constants.ServiceCodeProduct, constants.SuccessCode)
}

func (prod productsDelivery) RetrieveAllProducts(ctx *gin.Context) {
	productData, err := prod.productsUC.RetrieveAllProducts()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, productData, nil, "Success Retrieve All Products", constants.ServiceCodeProduct, constants.SuccessCode)
}

func (prod productsDelivery) RetrieveProductsByID(ctx *gin.Context) {
	productsID := ctx.Param("id")

	productData, err := prod.productsUC.RetrieveProductByID(productsID)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "product doesn't exist", constants.ServiceCodeProduct, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, productData, nil, "Success Retrieve Product By ID", constants.ServiceCodeProduct, constants.SuccessCode)
}

func (prod productsDelivery) UpdateProducts(ctx *gin.Context) {
	var req productsDto.ProductsRepo
	productIDStr := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	productsID, err := uuid.Parse(productIDStr)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	productsData := productsDto.ProductsRepo{
		ProductsID:   productsID,
		ProductName:  req.ProductName,
		ProductImage: req.ProductImage,
		Description:  req.Description,
		Price:        req.Price,
		CategoryID:   req.CategoryID,
		Stock:        req.Stock,
	}

	err = prod.productsUC.UpdateProducts(productsData)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "product doesn't exist", constants.ServiceCodeProduct, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "Success Update Product", constants.ServiceCodeProduct, constants.SuccessCode)
}

func (prod productsDelivery) DeleteProducts(ctx *gin.Context) {
	productIDStr := ctx.Param("id")

	err := prod.productsUC.DeleteProducts(productIDStr)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "product doesn't exist", constants.ServiceCodeProduct, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "Success Delete Product", constants.ServiceCodeProduct, constants.SuccessCode)
}
