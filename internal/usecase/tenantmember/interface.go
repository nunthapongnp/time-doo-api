package tenantmember

import (
	"time-doo-api/internal/domain"
)

type TenantMemberUsecase interface {
	GetTenantMemberByTenantID(tenantID int64) ([]*domain.TenantMember, error)
	FindTenantMemberByUserID(userID int64) (*domain.TenantMember, error)
	AddTenantMember(member *domain.TenantMember) error
	EditTenantMemberRole(id int64, role string) error
	RemoveTenantMember(id int64) error
}
