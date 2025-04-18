package project

import (
	"time-doo-api/internal/domain"
)

type ProjectRepository interface {
	Add(project *domain.Project) error
	FindByID(id int64) (*domain.Project, error)
	GetByTenantID(tenantID int64) ([]*domain.Project, error)
	Edit(project *domain.Project) error
	Remove(id int64) error
}
