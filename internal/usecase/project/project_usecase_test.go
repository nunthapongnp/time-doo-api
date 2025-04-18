package project_test

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"
// 	"time-doo-api/internal/domain"
// 	"time-doo-api/internal/usecase/project"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestProjectUsecase_Create(t *testing.T) {
// 	mockRepo := new(mockRepo.ProjectRepository)
// 	uc := project.NewProjectUsecase(mockRepo)

// 	project := &domain.Project{
// 		TenantID:  1,
// 		Name:      "Test Project",
// 		CreatedBy: 99,
// 		CreatedAt: time.Now(),
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockRepo.On("Create", mock.Anything, project).Return(nil).Once()
// 		err := uc.Create(context.Background(), project)
// 		assert.NoError(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("db error", func(t *testing.T) {
// 		mockRepo.On("Create", mock.Anything, project).Return(errors.New("db error")).Once()
// 		err := uc.Create(context.Background(), project)
// 		assert.EqualError(t, err, "db error")
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestProjectUsecase_GetByID(t *testing.T) {
// 	mockRepo := new(mockRepo.ProjectRepository)
// 	uc := project.NewProjectUsecase(mockRepo)

// 	project := &domain.Project{
// 		ID:       100,
// 		Name:     "Edge Case Project",
// 		TenantID: 1,
// 	}

// 	t.Run("project found", func(t *testing.T) {
// 		mockRepo.On("GetByID", mock.Anything, int64(100)).Return(project, nil).Once()
// 		p, err := uc.GetByID(context.Background(), 100)
// 		assert.NoError(t, err)
// 		assert.Equal(t, project, p)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("not found", func(t *testing.T) {
// 		mockRepo.On("GetByID", mock.Anything, int64(999)).Return(nil, errors.New("not found")).Once()
// 		p, err := uc.GetByID(context.Background(), 999)
// 		assert.Nil(t, p)
// 		assert.EqualError(t, err, "not found")
// 		mockRepo.AssertExpectations(t)
// 	})
// }
