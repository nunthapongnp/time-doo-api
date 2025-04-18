package task

import "time-doo-api/internal/domain"

type TaskUsecase interface {
	AddTask(t *domain.Task) error
	EditTask(t *domain.Task) error
	RemoveTask(id int64) error
	GetTaskByColumn(columnID int64) ([]*domain.Task, error)
	GetTaskByID(id int64) (*domain.Task, error)
	ReorderTask(columnID int64, ids []int64) error
	MoveTaskToColumn(taskID, fromColID, toColID int64, orderedTaskIDs []int64) error
}
