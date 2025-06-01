package user

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/model"
)

type UserUsecase interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByTenant(tenantID int64) ([]*model.UserDTO, error)
	AddUser(user *model.UserDTO) (*domain.User, error)
	EditUser(user *domain.User) error
	RemoveUser(id int64) error
}
