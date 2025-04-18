package user

import (
	"time-doo-api/internal/domain"
)

type UserUsecase interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	AddUser(user *domain.User) error
	EditUser(user *domain.User) error
	RemoveUser(id int64) error
}
