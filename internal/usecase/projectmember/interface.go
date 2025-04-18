package projectmember

import "time-doo-api/internal/domain"

type ProjectMemberUsecase interface {
	AddProjectMember(member *domain.ProjectMember) error
	RemoveProjectMember(projectID, userID int64) error
}
