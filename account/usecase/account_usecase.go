package usecase

import (
	"sample-api/domain"

	"github.com/gofrs/uuid"
)

type accountUsecase struct {
	accountRepo domain.AccountRepository
}

// NewAccountUsecase mount account usecase
func NewAccountUsecase(a domain.AccountRepository) domain.AccountUsecase {
	return &accountUsecase{
		accountRepo: a,
	}
}

func (a *accountUsecase) Get(id uuid.UUID) (*domain.User, error) {
	account, err := a.accountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:    account.ID,
		Email: account.Email,
	}
	return user, nil
}
