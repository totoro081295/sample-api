package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Account struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey" sql:"type:uuid" name:"ID"`
	Email     string     `json:"email" name:"email"`
	Password  string     `json:"password" name:"password"`
	CreatedBy string     `json:"createdBy" name:"作成者"`
	UpdatedBy string     `json:"updatedBy" name:"更新者"`
	CreatedAt time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time `json:"deletedAt" name:"削除日"`
}

type User struct {
	ID    uuid.UUID `json:"id" gorm:"primaryKey" sql:"type:uuid" name:"ID"`
	Email string    `json:"email" name:"email"`
}

// AccountRepository repository interface
type AccountRepository interface {
	GetByID(id uuid.UUID) (*Account, error)
	GetByEmail(email string) (*Account, error)
}

// AccountUsecase usecase interface
type AccountUsecase interface {
	Get(id uuid.UUID) (*User, error)
}
