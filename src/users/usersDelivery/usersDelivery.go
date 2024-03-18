package usersDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/usersDto"
	"E-Commerce/pkg/middleware"
	"E-Commerce/src/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

type userDelivery struct {
	userUC users.UserUseCase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC users.UserUseCase) {
	handler := userDelivery{
		userUC: userUC,
	}

	userGroup := v1Group.Group("/users")
	{
		userGroup.GET("", middleware.JWTAuth("admin"), handler.RetrieveAllUsers)
		userGroup.GET("/:id", middleware.JWTAuth("admin", "users"), handler.RetrieveUsersByID)
		userGroup.PUT("/:id", middleware.JWTAuth("admin", "users"), handler.UpdateProfiles)
	}
}

func (u userDelivery) RetrieveAllUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("size"))

	userData, pagination, err := u.userUC.RetrieveAllUsers(page, pageSize)
	if err != nil {
		if err.Error() == "01" {
			errorMessage := fmt.Sprintf("page %d doesn't exist", page)
			json.NewResponseForbidden(ctx, errorMessage, constants.ServiceCodeUsers, constants.Forbidden)
			return
		}

		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeUsers, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, userData, pagination, "success retrieve all users", constants.ServiceCodeUsers, constants.SuccessCode)
}

func (u userDelivery) RetrieveUsersByID(ctx *gin.Context) {
	usrID := ctx.Param("id")

	userData, err := u.userUC.RetrieveUsersByID(usrID)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "User profile is incomplete. Please update your profile before accessing detailed user information.", constants.ServiceCodeUsers, constants.Forbidden)
			return
		}

		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeUsers, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, userData, nil, "success retrieve users", constants.ServiceCodeUsers, constants.SuccessCode)
}

func (u userDelivery) UpdateProfiles(ctx *gin.Context) {
	var req usersDto.UserProfile
	usrIDStr := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeUsers, constants.GeneralErrCode)
		return
	}

	usrID, err := uuid.Parse(usrIDStr)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeUsers, constants.GeneralErrCode)
		return
	}

	usrProfile := usersDto.UserUpdate{
		UserID:     usrID,
		FullName:   req.FullName,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
		Phone:      req.Phone,
	}

	err = u.userUC.UpdateProfiles(usrProfile)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeUsers, constants.GeneralErrCode)
		return
	}

	usrData, err := u.userUC.RetrieveUsersByID(usrIDStr)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeUsers, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, usrData, nil, "success retrieve users", constants.ServiceCodeUsers, constants.SuccessCode)
}
