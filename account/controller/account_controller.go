package controller

import (
	"log"
	"net/http"
	"sample-api/domain"
	"sample-api/utils/status"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// AccountController account controller
type AccountController struct {
	AccountUsecase domain.AccountUsecase
	Token          domain.TokenHandler
}

// NewAccountController mount account controller
func NewAccountController(e *echo.Echo, accountUcase domain.AccountUsecase, token domain.TokenHandler) {
	handler := &AccountController{
		AccountUsecase: accountUcase,
		Token:          token,
	}
	// external api group
	api := e.Group("/api", middleware.JWT([]byte("secret")))
	api.GET("/accounts", handler.GetByToken)
}

// Get get account
func (c *AccountController) GetByToken(ctx echo.Context) error {
	id, err := c.Token.GetToken(ctx)
	if err != nil {
		log.Println(err)
		if err == status.ErrUnauthorized {
			res := status.ResStatus{
				Result:  "failure",
				Message: "Could not authenticate with the requested jwt.",
			}
			return ctx.JSON(http.StatusUnauthorized, res)
		}
		return status.ResponseError(ctx, err)
	}
	a, err := c.AccountUsecase.Get(id)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, a)
}
