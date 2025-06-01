package tenantmember

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewTenantMemberRepository(db *ctx.AppDbContext) TenantMemberRepository {
	return &repository{db: db}
}

func (r *repository) GetByTenantID(tenantID int64) ([]*domain.TenantMember, error) {
	var members []*domain.TenantMember
	err := r.db.Raw().Preload("Users").Preload("Tenants").Find(&members, "tenant_id = ?", tenantID).Error
	return members, err
}

func (r *repository) FindByUserID(userID int64) (*domain.TenantMember, error) {
	var member domain.TenantMember
	err := r.db.Raw().Preload("Users").Preload("Tenants").First(&member, "user_id = ?", userID).Error
	return &member, err
}

func (r *repository) Add(member *domain.TenantMember) error {
	return r.db.Create(member)
}

func (r *repository) EditRole(id int64, role string) error {
	return r.db.Raw().Model(&domain.TenantMember{}).Where("id = ?", id).Update("role", role).Error
}

func (r *repository) Reomve(id int64) error {
	return r.db.Delete(&domain.TenantMember{ID: id})
}
