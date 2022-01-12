package repository

import (
	"errors"
	"sample-api/domain"
	"sample-api/utils/status"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type accountRepository struct {
	Conn *gorm.DB
}

// NewAccountRepository mount account repository
func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return &accountRepository{
		Conn: db,
	}
}

func (a *accountRepository) GetByID(id uuid.UUID) (*domain.Account, error) {
	account := &domain.Account{}
	err := a.Conn.Model(&account).Where("id = ?", id).Find(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.ErrNotFound
	}
	return account, nil
}

func (a *accountRepository) GetByEmail(email string) (*domain.Account, error) {
	account := &domain.Account{}
	err := a.Conn.Model(&account).Where("email = ?", email).Find(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.ErrNotFound
	}
	return account, nil
}
