package usersDelivery

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/json"
	"E-Commerce/pkg/middleware"
	"E-Commerce/src/users"
	"fmt"
	"github.com/gin-gonic/gin"
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
