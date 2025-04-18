package task

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/task"
)

type usecase struct {
	taskRepo task.TaskRepository
}

func NewTaskUsecase(taskRepo task.TaskRepository) TaskUsecase {
	return &usecase{taskRepo}
}

func (u *usecase) AddTask(t *domain.Task) error {
	return u.taskRepo.Add(t)
}

func (u *usecase) EditTask(t *domain.Task) error {
	return u.taskRepo.Edit(t)
}

func (u *usecase) RemoveTask(id int64) error {
	return u.taskRepo.Reomve(id)
}

func (u *usecase) GetTaskByColumn(columnID int64) ([]*domain.Task, error) {
	return u.taskRepo.GetByColumn(columnID)
}

func (u *usecase) GetTaskByID(id int64) (*domain.Task, error) {
	return u.taskRepo.FindByID(id)
}

func (u *usecase) ReorderTask(columnID int64, ids []int64) error {
	return u.taskRepo.Reorder(columnID, ids)
}

func (u *usecase) MoveTaskToColumn(taskID, fromColID, toColID int64, orderedTaskIDs []int64) error {
	return u.taskRepo.MoveTask(taskID, fromColID, toColID, orderedTaskIDs)
}
