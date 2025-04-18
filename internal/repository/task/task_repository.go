package task

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"

	"gorm.io/gorm"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewTaskRepository(db *ctx.AppDbContext) TaskRepository {
	return &repository{db: db}
}

func (r *repository) Add(t *domain.Task) error {
	return r.db.Create(t)
}

func (r *repository) Edit(t *domain.Task) error {
	return r.db.Update(t, t)
}

func (r *repository) Reomve(id int64) error {
	return r.db.Delete(&domain.Task{ID: id})
}

func (r *repository) GetByColumn(columnID int64) ([]*domain.Task, error) {
	var tasks []*domain.Task
	err := r.db.Raw().Where("column_id = ?", columnID).Order("position ASC").Find(&tasks).Error
	return tasks, err
}

func (r *repository) FindByID(id int64) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *repository) Reorder(columnID int64, ids []int64) error {
	return r.db.Raw().Transaction(func(tx *gorm.DB) error {
		for i, id := range ids {
			if err := tx.Model(&domain.Task{}).
				Where("id = ? AND column_id = ?", id, columnID).
				Update("position", i).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *repository) MoveTask(taskID, fromColID, toColID int64, orderedIDs []int64) error {
	return r.db.Raw().Transaction(func(tx *gorm.DB) error {
		// Move task to new column
		if err := tx.Model(&domain.Task{}).
			Where("id = ?", taskID).
			Updates(map[string]interface{}{
				"column_id": toColID,
			}).Error; err != nil {
			return err
		}

		// Reorder tasks in new column
		for i, id := range orderedIDs {
			if err := tx.Model(&domain.Task{}).
				Where("id = ? AND column_id = ?", id, toColID).
				Update("position", i).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
