package projectmember

import "time-doo-api/internal/domain"

type ProjectMemberRepository interface {
	Add(member *domain.ProjectMember) error
	GetByProjectID(projectID int64) ([]*domain.ProjectMember, error)
	Remove(projectID, userID int64) error
}
