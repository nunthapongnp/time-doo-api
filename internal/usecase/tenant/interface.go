package tenant

import (
	"time-doo-api/internal/domain"
)

type TenantUsecase interface {
	GetTenantMembers(tenantID int64) ([]*domain.TenantMember, error)
	AddTenant(tenant *domain.Tenant) error
	GetTenantByID(id int64) (*domain.Tenant, error)
	GetAllTenants() ([]*domain.Tenant, error)
}
