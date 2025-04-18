package tenant

import (
	"time-doo-api/internal/domain"
)

type TenantRepository interface {
	Add(tenant *domain.Tenant) error
	FindByID(id int64) (*domain.Tenant, error)
	GetAll() ([]*domain.Tenant, error)
}
