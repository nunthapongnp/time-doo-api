package project

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/project"
	"time-doo-api/internal/repository/projectmember"
)

type usecase struct {
	projectRepo       project.ProjectRepository
	projectmemberRepo projectmember.ProjectMemberRepository
}

func NewProjectUsecase(projectRepo project.ProjectRepository, projectmemberRepo projectmember.ProjectMemberRepository) ProjectUsecase {
	return &usecase{projectRepo, projectmemberRepo}
}

func (uc *usecase) AddProject(p *domain.Project) error {
	return uc.projectRepo.Add(p)
}

func (uc *usecase) GetProjectByID(id int64) (*domain.Project, error) {
	return uc.projectRepo.FindByID(id)
}

func (uc *usecase) GetProjectByTenant(tenantID int64) ([]*domain.Project, error) {
	return uc.projectRepo.GetByTenantID(tenantID)
}

func (uc *usecase) EditProject(p *domain.Project) error {
	return uc.projectRepo.Edit(p)
}

func (uc *usecase) RemoveProject(id int64) error {
	return uc.projectRepo.Remove(id)
}

func (u *usecase) GetProjectMembers(projectID int64) ([]*domain.ProjectMember, error) {
	return u.projectmemberRepo.GetByProjectID(projectID)
}
