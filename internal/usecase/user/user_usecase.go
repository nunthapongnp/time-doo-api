package user

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/user"
)

type usecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) UserUsecase {
	return &usecase{userRepo}
}

func (u *usecase) GetAllUsers() ([]*domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *usecase) GetUserByID(id int64) (*domain.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *usecase) GetUserByEmail(email string) (*domain.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *usecase) AddUser(user *domain.User) error {
	return u.userRepo.Add(user)
}

func (u *usecase) EditUser(user *domain.User) error {
	return u.userRepo.Edit(user)
}

func (u *usecase) RemoveUser(id int64) error {
	return u.userRepo.Remove(id)
}
