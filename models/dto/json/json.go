package json

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type (
	// JSONResponse - struct for json response success
	jsonResponse struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
		Paging  interface{} `json:"paging,omitempty"`
	}

	// JSONResponse - struct for json response success
	jsonErrorResponse struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

	ValidationField struct {
		FieldName string `json:"field"`
		Message   string `json:"message"`
	}

	// JSONResponse - struct for json response success
	jsonBadRequestResponse struct {
		Code             string            `json:"responseCode"`
		Message          string            `json:"responseMessage"`
		ErrorDescription []ValidationField `json:"error_description,omitempty"`
	}

	Pagination struct {
		CurrentPage  int `json:"page"`
		TotalPages   int `json:"totalPages"`
		TotalRecords int `json:"totalData"`
	}
)

func NewResponseSuccess(c *gin.Context, result interface{}, pagination interface{}, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    "200" + serviceCode + responseCode,
		Message: message,
		Data:    result,
		Paging:  pagination,
	})
}

func NewResponseBadRequest(c *gin.Context, validationField []ValidationField, message, serviceCode, errorCode string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Code:             "400" + serviceCode + errorCode,
		Message:          message,
		ErrorDescription: validationField,
	})
}

func NewResponseError(c *gin.Context, err, serviceCode, errorCode string) {
	log.Error().Msg(err)
	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Code:    "500" + serviceCode + errorCode,
		Message: "internal server error",
		Error:   err,
	})
}

func NewResponseForbidden(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401" + serviceCode + errorCode,
		Message: message,
	})
}
