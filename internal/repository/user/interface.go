package user

import (
	"time-doo-api/internal/domain"
)

type UserRepository interface {
	Add(user *domain.User) error
	FindByID(id int64) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	GetByTenant(tenantID int64) ([]*domain.User, error)
	Edit(user *domain.User) error
	Remove(id int64) error
}
