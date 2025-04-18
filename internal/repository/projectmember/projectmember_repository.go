package projectmember

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewProjectMemberRepository(db *ctx.AppDbContext) ProjectMemberRepository {
	return &repository{db: db}
}

func (r *repository) Add(member *domain.ProjectMember) error {
	return r.db.Create(member)
}

func (r *repository) GetByProjectID(projectID int64) ([]*domain.ProjectMember, error) {
	var members []*domain.ProjectMember
	err := r.db.Find(&members, "project_id = ?", projectID)
	return members, err
}

func (r *repository) Remove(projectID, userID int64) error {
	return r.db.Raw().Where("project_id = ? AND user_id = ?", projectID, userID).Delete(&domain.ProjectMember{}).Error
}
