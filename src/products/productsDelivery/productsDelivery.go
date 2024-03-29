package productsDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/productsDto"
	"E-Commerce/pkg/validation"
	"E-Commerce/src/products"
	"github.com/gin-gonic/gin"
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
	defer fileContent.Close()

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
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeProduct, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, product, nil, "Success Create Product", constants.ServiceCodeProduct, constants.SuccessCode)
}
