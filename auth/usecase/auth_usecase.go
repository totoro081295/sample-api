package usecase

import (
	"sample-api/domain"
	"sample-api/utils"
	"sample-api/utils/status"
)

type authUsecase struct {
	accountRepo domain.AccountRepository
	token       domain.TokenHandler
}

func NewAuthUsecase(a domain.AccountRepository, t domain.TokenHandler) domain.AuthUsecase {
	return &authUsecase{
		accountRepo: a,
		token:       t,
	}
}

func (a *authUsecase) Login(l *domain.Auth) (*domain.Token, error) {
	u, err := a.accountRepo.GetByEmail(l.Email)
	if err != nil {
		return nil, err
	}
	err = utils.CompareHashPassword(u.Password, l.Password)
	if err != nil {
		return nil, status.ErrUnauthorized
	}

	token, err := a.token.GenerateJWT(u.ID, false)
	if err != nil {
		return nil, err
	}

	res := domain.Token{
		AccessToken: token,
	}
	return &res, nil
}
