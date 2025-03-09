package services

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nunthapongnp/time-doo-api/internal/models"
	"github.com/nunthapongnp/time-doo-api/internal/repositories"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// Task methods
func (s *TaskService) CreateTask(ctx *gin.Context, task *models.Task) (string, error) {
	task.CreatedBy = ctx.GetString("userID")
	task.CreatedDate = time.Now()
	task.RowVersion = int(time.Now().Unix())
	return s.repo.CreateTask(ctx, task)
}

func (s *TaskService) GetTask(ctx *gin.Context, id string) (*models.Task, error) {
	return s.repo.GetTask(ctx, id)
}

func (s *TaskService) UpdateTask(ctx *gin.Context, id string, task *models.Task) error {
	// Preserve original created date
	existingTask, err := s.repo.GetTask(ctx, id)
	if err != nil {
		return err
	}

	if existingTask == nil {
		return fmt.Errorf("task not found")
	}

	if task.RowVersion != existingTask.RowVersion {
		return fmt.Errorf("task has been updated by another user")
	}

	task.UpdatedBy = ctx.GetString("userID")
	return s.repo.UpdateTask(ctx, id, task)
}

func (s *TaskService) DeleteTask(ctx *gin.Context, id string) error {
	return s.repo.DeleteTask(ctx, id)
}
