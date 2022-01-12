package utils

import (
	"sample-api/domain"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

type tokenHandler struct{}

// NewTokenHandler mount token handler
func NewTokenHandler() domain.TokenHandler {
	return &tokenHandler{}
}

type JWTCustomClaims struct {
	jwt.StandardClaims
}

// GenerateJWT トークン生成
func (h *tokenHandler) GenerateJWT(id uuid.UUID, valid bool) (string, error) {
	// Set claims
	claims := JWTCustomClaims{
		jwt.StandardClaims{
			Issuer:    "sample-api",
			Subject:   id.String(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(1800)).Unix(), // 30分
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *tokenHandler) GetToken(ctx echo.Context) (uuid.UUID, error) {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	accountID, err := uuid.FromString(sub)
	if err != nil {
		return uuid.Nil, err
	}

	return accountID, nil
}
