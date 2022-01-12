package domain

import "sample-api/utils/status"

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

// AuthUsecase auth usecase interface
type AuthUsecase interface {
	Login(l *Auth) (*Token, error)
}

// LoginResponse login response for external
type LoginResponse struct {
	Status      status.ResStatus `json:"status"`
	AccessToken string           `json:"accessToken"`
}
