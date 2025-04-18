package tenantmember

import (
	"time-doo-api/internal/domain"
)

type TenantMemberUsecase interface {
	AddTenantMember(member *domain.TenantMember) error
	EditTenantMemberRole(id int64, role string) error
	RemoveTenantMember(id int64) error
}
