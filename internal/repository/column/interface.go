package column

import "time-doo-api/internal/domain"

type ColumnRepository interface {
	Add(c *domain.Column) error
	GetByProject(projectID int64) ([]*domain.Column, error)
	Edit(c *domain.Column) error
	Remove(id int64) error
	Reorder(projectID int64, ids []int64) error
	GetWithTasks(projectID int64) ([]*domain.Column, error)
}
