package authenticationDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/authenticationDto"
	"E-Commerce/models/dto/json"
	"E-Commerce/pkg/validation"
	"E-Commerce/src/authentication"
	"github.com/gin-gonic/gin"
)

type authenticationDelivery struct {
	authenticationUC authentication.AuthenticationUseCase
}

func NewAuthenticationDelivery(v1Group *gin.RouterGroup, authenticationUC authentication.AuthenticationUseCase) {
	handler := authenticationDelivery{
		authenticationUC: authenticationUC,
	}

	authenticationGroup := v1Group.Group("/auth")
	{
		authenticationGroup.POST("/register", handler.Register)
		authenticationGroup.POST("/login", handler.Login)
	}
}

func (auth authenticationDelivery) Register(ctx *gin.Context) {
	var req authenticationDto.RegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeAuth, constants.GeneralErrCode)
		return
	}
	validationErr := validation.ValidateRegister(req)
	if len(validationErr) > 0 {
		json.NewResponseBadRequest(ctx, validationErr, constants.BadReqMsg, constants.ServiceCodeAuth, constants.GeneralErrCode)
		return
	}
	usrReq, err := auth.authenticationUC.RegisterUsers(req)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "email already registered", constants.ServiceCodeAuth, constants.Forbidden)
			return
		}
		if err.Error() == "02" {
			json.NewResponseForbidden(ctx, "username already registered", constants.ServiceCodeAuth, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeAuth, constants.GeneralErrCode)
		return
	}
	json.NewResponseSuccess(ctx, usrReq, nil, "user registered successfully.", constants.ServiceCodeAuth, constants.SuccessCode)
}

func (auth authenticationDelivery) Login(ctx *gin.Context) {
	var req authenticationDto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeAuth, constants.GeneralErrCode)
		return
	}

	validationErr := validation.ValidateLogin(req)
	if len(validationErr) > 0 {
		json.NewResponseBadRequest(ctx, validationErr, constants.BadReqMsg, constants.ServiceCodeAuth, constants.GeneralErrCode)
		return
	}

	token, err := auth.authenticationUC.LoginUsers(req)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "email doesn't exists on our records", constants.ServiceCodeAuth, constants.Forbidden)
			return
		}
		if err.Error() == "02" {
			json.NewResponseForbidden(ctx, "Unauthorized email and password didn't match", constants.ServiceCodeAuth, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeAuth, constants.GeneralErrCode)
		return
	}
	data := interface{}(map[string]string{"access_token": token})
	json.NewResponseSuccess(ctx, data, nil, "login successfully.", constants.ServiceCodeAuth, constants.SuccessCode)
}
