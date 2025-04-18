package tenant

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/tenant"
	"time-doo-api/internal/repository/tenantmember"
)

type usecase struct {
	tenantRepo       tenant.TenantRepository
	tenantMemberRepo tenantmember.TenantMemberRepository
}

func NewTenantUsecase(tenantRepo tenant.TenantRepository, tenantMemberRepo tenantmember.TenantMemberRepository) TenantUsecase {
	return &usecase{tenantRepo, tenantMemberRepo}
}

func (u *usecase) GetTenantMembers(tenantID int64) ([]*domain.TenantMember, error) {
	return u.tenantMemberRepo.GetByTenantID(tenantID)
}

func (u *usecase) AddTenant(tenant *domain.Tenant) error {
	return u.tenantRepo.Add(tenant)
}

func (u *usecase) GetTenantByID(id int64) (*domain.Tenant, error) {
	return u.tenantRepo.FindByID(id)
}

func (u *usecase) GetAllTenants() ([]*domain.Tenant, error) {
	return u.tenantRepo.GetAll()
}
