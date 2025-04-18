package project

import (
	"time-doo-api/internal/domain"
)

type ProjectUsecase interface {
	AddProject(p *domain.Project) error
	GetProjectByID(id int64) (*domain.Project, error)
	GetProjectByTenant(tenantID int64) ([]*domain.Project, error)
	EditProject(p *domain.Project) error
	RemoveProject(id int64) error
	GetProjectMembers(projectID int64) ([]*domain.ProjectMember, error)
}
