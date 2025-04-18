package tenant

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewTenantRepository(db *ctx.AppDbContext) TenantRepository {
	return &repository{db: db}
}

func (r *repository) Add(tenant *domain.Tenant) error {
	return r.db.Create(tenant)
}

func (r *repository) FindByID(id int64) (*domain.Tenant, error) {
	var tenant domain.Tenant
	err := r.db.Find(&tenant, id)
	return &tenant, err
}

func (r *repository) GetAll() ([]*domain.Tenant, error) {
	var tenants []*domain.Tenant
	err := r.db.Raw().Order("id DESC").Find(&tenants).Error
	return tenants, err
}
