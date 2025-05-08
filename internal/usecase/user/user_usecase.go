package user

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/tenantmember"
	"time-doo-api/internal/repository/user"
)

type usecase struct {
	userRepo         user.UserRepository
	tenantMemberRepo tenantmember.TenantMemberRepository
}

func NewUserUsecase(userRepo user.UserRepository, tenantMemberRepo tenantmember.TenantMemberRepository) UserUsecase {
	return &usecase{userRepo, tenantMemberRepo}
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
	if err := u.userRepo.Add(user); err != nil {
		return err
	}

	member := &domain.TenantMember{
		TenantID: user.TenantID,
		UserID:   user.ID,
		Role:     "admin",
	}

	return u.tenantMemberRepo.Add(member)
}

func (u *usecase) EditUser(user *domain.User) error {
	return u.userRepo.Edit(user)
}

func (u *usecase) RemoveUser(id int64) error {
	return u.userRepo.Remove(id)
}

func (u *usecase) GetUserByTenant(tenantID int64) ([]*domain.User, error) {
	return u.userRepo.GetByTenant(tenantID)
}
