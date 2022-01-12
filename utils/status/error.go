package status

import (
	"net/http"

	"github.com/pkg/errors"

	echo "github.com/labstack/echo/v4"
)

var (
	// ErrInternalServer internal server error
	ErrInternalServer = errors.New("Internal server error")
	// ErrNotFound not found error
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict conflict error
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadRequest bad request error
	ErrBadRequest = errors.New("Bad request")
	// ErrUnauthorized unauthorized error
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrForbidden forbidden error
	ErrForbidden = errors.New("Forbidden")
)

// ErrorMessage error message
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResStatus response status struct
type ResStatus struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

// ErrResponse error response
type ErrResponse struct {
	Status ResStatus `json:"status"`
}

// ResponseError 返却するエラーコードの指定
func ResponseError(ctx echo.Context, err error) error {
	res := ErrorMessage{}
	switch errors.Cause(err) {
	case ErrInternalServer:
		res.Code = 500
		res.Message = err.Error()
		return ctx.JSON(500, res)
	case ErrNotFound:
		res.Code = 404
		res.Message = err.Error()
		return ctx.JSON(http.StatusNotFound, res)
	case ErrConflict:
		res.Code = http.StatusConflict
		res.Message = err.Error()
		return ctx.JSON(http.StatusConflict, res)
	case ErrBadRequest:
		res.Code = 400
		res.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, res)
	case ErrUnauthorized:
		res.Code = http.StatusUnauthorized
		res.Message = err.Error()
		return ctx.JSON(http.StatusUnauthorized, res)
	case ErrForbidden:
		res.Code = http.StatusForbidden
		res.Message = err.Error()
		return ctx.JSON(http.StatusForbidden, res)
	default:
		res.Code = 500
		res.Message = err.Error()
		return ctx.JSON(500, res)
	}
}
