package auth

import (
	"errors"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/tenant"
	"time-doo-api/internal/repository/tenantmember"
	"time-doo-api/internal/repository/user"
	pwd "time-doo-api/pkg/bcrypt"
	"time-doo-api/pkg/jwt"
)

type usecase struct {
	userRepo         user.UserRepository
	tenantRepo       tenant.TenantRepository
	tenantMemberRepo tenantmember.TenantMemberRepository
}

func NewAuthUsecase(userRepo user.UserRepository, tenantRepo tenant.TenantRepository, tenantMemberRepo tenantmember.TenantMemberRepository) AuthUsecase {
	return &usecase{userRepo, tenantRepo, tenantMemberRepo}
}

func (u *usecase) Register(tenantName string, user *domain.User) error {
	if tenantName == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing tenant name or user info")
	}

	tenant := &domain.Tenant{Name: tenantName}
	if err := u.tenantRepo.Add(tenant); err != nil {
		return err
	}

	user.TenantID = tenant.ID
	if err := u.userRepo.Add(user); err != nil {
		return err
	}

	member := &domain.TenantMember{
		TenantID: tenant.ID,
		UserID:   user.ID,
		Role:     "admin",
	}

	return u.tenantMemberRepo.Add(member)
}

func (u *usecase) Login(email, password string) (string, error) {
	usr, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("not found user with this email")
	}

	mem, err := u.tenantMemberRepo.FindByUserID(usr.ID)
	if err != nil {
		return "", errors.New("not found user in tenant")
	}

	if !pwd.VerifyPassword(usr.Password, password) {
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(uint(usr.ID), uint(mem.TenantID), mem.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
