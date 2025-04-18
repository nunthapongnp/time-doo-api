package projectmember

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/projectmember"
)

type usecase struct {
	projectmemberRepo projectmember.ProjectMemberRepository
}

func NewProjectMemberUsecase(projectmemberRepo projectmember.ProjectMemberRepository) ProjectMemberUsecase {
	return &usecase{projectmemberRepo}
}

func (u *usecase) AddProjectMember(member *domain.ProjectMember) error {
	return u.projectmemberRepo.Add(member)
}

func (u *usecase) RemoveProjectMember(projectID, userID int64) error {
	return u.projectmemberRepo.Remove(projectID, userID)
}
