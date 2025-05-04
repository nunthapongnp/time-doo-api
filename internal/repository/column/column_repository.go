package column

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"

	"gorm.io/gorm"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewColumnRepository(db *ctx.AppDbContext) ColumnRepository {
	return &repository{db: db}
}

func (r *repository) Add(c *domain.Column) error {
	return r.db.Create(c)
}

func (r *repository) GetByProject(projectID int64) ([]*domain.Column, error) {
	var cols []*domain.Column
	err := r.db.Raw().Where("project_id = ?", projectID).Order("position ASC").Find(&cols).Error
	return cols, err
}

func (r *repository) Edit(c *domain.Column) error {
	return r.db.Update(c, c)
}

func (r *repository) Remove(id int64) error {
	return r.db.Delete(&domain.Column{ID: id})
}

func (r *repository) Reorder(projectID int64, ids []int64) error {
	return r.db.Raw().Transaction(func(tx *gorm.DB) error {
		for i, id := range ids {
			if err := tx.Model(&domain.Column{}).
				Where("id = ? AND projectid = ?", id, projectID).
				Update("position", i).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *repository) GetWithTasks(projectID int64) ([]*domain.Column, error) {
	var cols []*domain.Column
	err := r.db.Raw().Where("project_id = ?", projectID).Order("position ASC").Preload("Tasks").Find(&cols).Error
	return cols, err
}
