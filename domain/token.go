package domain

import (
	"github.com/gofrs/uuid"
	echo "github.com/labstack/echo/v4"
)

// TokenHandler token handler interface
type TokenHandler interface {
	GenerateJWT(id uuid.UUID, valid bool) (string, error)
	GetToken(ctx echo.Context) (uuid.UUID, error)
}
