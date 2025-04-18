package auth

import (
	"time-doo-api/internal/domain"
)

type AuthUsecase interface {
	Register(tenantName string, user *domain.User) error
	Login(email, password string) (string, error)
}
