package project

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewProjectRepository(db *ctx.AppDbContext) ProjectRepository {
	return &repository{db: db}
}

func (r *repository) Add(p *domain.Project) error {
	return r.db.Create(p)
}

func (r *repository) FindByID(id int64) (*domain.Project, error) {
	var p domain.Project
	err := r.db.Find(&p, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) GetByTenantID(tenantID int64) ([]*domain.Project, error) {
	var projects []*domain.Project
	err := r.db.Raw().Where("tenant_id = ?", tenantID).Order("id DESC").Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *repository) Edit(p *domain.Project) error {
	return r.db.Update(p, p)
}

func (r *repository) Remove(id int64) error {
	return r.db.Delete(&domain.Project{ID: id})
}
