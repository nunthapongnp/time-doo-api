package column

import "time-doo-api/internal/domain"

type ColumnUsecase interface {
	AddColumn(c *domain.Column) error
	GetColumnByProject(projectID int64) ([]*domain.Column, error)
	EditColumn(c *domain.Column) error
	RemoveColumn(id int64) error
	ReorderColumn(projectID int64, ids []int64) error
	GetColumnWithTasks(projectID int64) ([]*domain.Column, error)
}
