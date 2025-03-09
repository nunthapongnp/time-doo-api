package services

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nunthapongnp/time-doo-api/internal/models"
	"github.com/nunthapongnp/time-doo-api/internal/repositories"
)

type SubTaskService struct {
	repo *repositories.SubTaskRepository
}

func NewSubTaskService(repo *repositories.SubTaskRepository) *SubTaskService {
	return &SubTaskService{repo: repo}
}

func (s *SubTaskService) CreateSubtask(ctx *gin.Context, taskID string, subtask *models.Subtask) (string, error) {
	subtask.CreatedBy = ctx.GetString("userID")
	subtask.CreatedDate = time.Now()
	subtask.RowVersion = int(time.Now().Unix())
	return s.repo.CreateSubtask(ctx, taskID, subtask)
}

func (s *SubTaskService) GetSubtask(ctx *gin.Context, taskID, subtaskID string) (*models.Subtask, error) {
	return s.repo.GetSubtask(ctx, taskID, subtaskID)
}

func (s *SubTaskService) UpdateSubtask(ctx *gin.Context, taskID, subtaskID string, subtask *models.Subtask) error {
	// Preserve original created date
	existingSubtask, err := s.repo.GetSubtask(ctx, taskID, subtaskID)
	if err != nil {
		return err
	}

	if existingSubtask == nil {
		return fmt.Errorf("subtask not found")
	}

	if subtask.RowVersion != existingSubtask.RowVersion {
		return fmt.Errorf("task has been updated by another user")
	}
	return s.repo.UpdateSubtask(ctx, taskID, subtaskID, subtask)
}

func (s *SubTaskService) DeleteSubtask(ctx *gin.Context, taskID, subtaskID string) error {
	return s.repo.DeleteSubtask(ctx, taskID, subtaskID)
}
