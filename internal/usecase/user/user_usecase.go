package user

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/model"
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

func (u *usecase) AddUser(model *model.UserDTO) (*domain.User, error) {
	user := &domain.User{
		FullName: model.FullName,
		Email:    model.Email,
		Password: model.Password,
	}

	if err := u.userRepo.Add(user); err != nil {
		return user, err
	}

	member := &domain.TenantMember{
		TenantID: model.TenantID,
		UserID:   user.ID,
		Role:     model.Role,
	}

	if err := u.tenantMemberRepo.Add(member); err != nil {
		return user, err
	}

	return user, nil
}

func (u *usecase) EditUser(user *domain.User) error {
	return u.userRepo.Edit(user)
}

func (u *usecase) RemoveUser(id int64) error {
	return u.userRepo.Remove(id)
}

func (u *usecase) GetUserByTenant(tenantID int64) ([]*model.UserDTO, error) {
	mem, err := u.tenantMemberRepo.GetByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	var users []*model.UserDTO
	for _, m := range mem {
		users = append(users, &model.UserDTO{
			ID:       m.UserID,
			Email:    m.Users.Email,
			FullName: m.Users.FullName,
			TenantID: m.TenantID,
			Role:     m.Role,
			AppModel: m.AppModel,
		})
	}

	return users, err
}
