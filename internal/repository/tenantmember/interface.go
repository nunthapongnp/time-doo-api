package tenantmember

import (
	"time-doo-api/internal/domain"
)

type TenantMemberRepository interface {
	GetByTenantID(tenantID int64) ([]*domain.TenantMember, error)
	FindByUserID(userID int64) (*domain.TenantMember, error)
	Add(member *domain.TenantMember) error
	EditRole(id int64, role string) error
	Reomve(id int64) error
}
