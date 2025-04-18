package tenantmember

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/tenantmember"
)

type usecase struct {
	tenantMemberRepo tenantmember.TenantMemberRepository
}

func NewTenantMemberUsecase(tenantMemberRepo tenantmember.TenantMemberRepository) TenantMemberUsecase {
	return &usecase{tenantMemberRepo}
}

func (u *usecase) AddTenantMember(member *domain.TenantMember) error {
	return u.tenantMemberRepo.Add(member)
}

func (u *usecase) EditTenantMemberRole(id int64, role string) error {
	return u.tenantMemberRepo.EditRole(id, role)
}

func (u *usecase) RemoveTenantMember(id int64) error {
	return u.tenantMemberRepo.Reomve(id)
}
