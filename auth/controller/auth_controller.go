package controller

import (
	"log"
	"net/http"
	"sample-api/domain"
	"sample-api/utils/status"

	"github.com/labstack/echo/v4"
)

// AuthController auth controller
type AuthController struct {
	AuthUsecase domain.AuthUsecase
}

// NewAuthController mount auth controller
func NewAuthController(e *echo.Echo, authUcase domain.AuthUsecase) {
	handler := &AuthController{
		AuthUsecase: authUcase,
	}
	e.POST("/oauth/login", handler.Login)
}

func (c *AuthController) Login(ctx echo.Context) error {
	request := domain.Auth{}
	err := ctx.Bind(&request)
	if err != nil {
		resStatus := status.ResStatus{
			Result:  "failure",
			Message: err.Error(),
		}
		res := status.ErrResponse{
			Status: resStatus,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, res)
	}

	// validation
	err = domain.LoginValidate(&request)
	if err != nil {
		log.Println(err)
		msg := "There is some errors in the requested data."
		resStatus := status.ResStatus{
			Result:  "failure",
			Message: msg,
		}
		res := status.ErrResponse{
			Status: resStatus,
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	token, err := c.AuthUsecase.Login(&request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	resStatus := status.ResStatus{
		Result:  "success",
		Message: "login is success.",
	}
	res := domain.LoginResponse{
		Status:      resStatus,
		AccessToken: token.AccessToken,
	}
	return ctx.JSON(http.StatusOK, res)
}
