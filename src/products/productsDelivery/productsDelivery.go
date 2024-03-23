package productsDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
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
