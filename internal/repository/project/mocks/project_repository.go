package mocks

// import (
// 	"context"
// 	"time-doo-api/internal/domain"

// 	"github.com/stretchr/testify/mock"
// )

// type ProjectRepository struct {
// 	mock.Mock
// }

// func (m *ProjectRepository) Create(ctx context.Context, p *domain.Project) error {
// 	args := m.Called(ctx, p)
// 	return args.Error(0)
// }

// func (m *ProjectRepository) GetByID(ctx context.Context, id int64) (*domain.Project, error) {
// 	args := m.Called(ctx, id)
// 	return args.Get(0).(*domain.Project), args.Error(1)
// }

// func (m *ProjectRepository) GetByTenantID(ctx context.Context, tenantID int64) ([]*domain.Project, error) {
// 	args := m.Called(ctx, tenantID)
// 	return args.Get(0).([]*domain.Project), args.Error(1)
// }

// func (m *ProjectRepository) Update(ctx context.Context, p *domain.Project) error {
// 	args := m.Called(ctx, p)
// 	return args.Error(0)
// }

// func (m *ProjectRepository) Delete(ctx context.Context, id int64) error {
// 	args := m.Called(ctx, id)
// 	return args.Error(0)
// }
