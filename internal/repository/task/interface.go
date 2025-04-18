package task

import "time-doo-api/internal/domain"

type TaskRepository interface {
	Add(t *domain.Task) error
	Edit(t *domain.Task) error
	Reomve(id int64) error
	GetByColumn(columnID int64) ([]*domain.Task, error)
	FindByID(id int64) (*domain.Task, error)
	Reorder(columnID int64, ids []int64) error
	MoveTask(taskID, fromColID, toColID int64, orderedIDs []int64) error
}
