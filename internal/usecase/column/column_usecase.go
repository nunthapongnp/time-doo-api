package column

import (
	"time-doo-api/internal/domain"
	"time-doo-api/internal/repository/column"
)

type usecase struct {
	columnRepo column.ColumnRepository
}

func NewColumnUsecase(columnRepo column.ColumnRepository) ColumnUsecase {
	return &usecase{columnRepo}
}

func (u *usecase) AddColumn(c *domain.Column) error {
	return u.columnRepo.Add(c)
}

func (u *usecase) GetColumnByProject(projectID int64) ([]*domain.Column, error) {
	return u.columnRepo.GetByProject(projectID)
}

func (u *usecase) EditColumn(c *domain.Column) error {
	return u.columnRepo.Edit(c)
}

func (u *usecase) RemoveColumn(id int64) error {
	return u.columnRepo.Remove(id)
}

func (u *usecase) ReorderColumn(projectID int64, ids []int64) error {
	return u.columnRepo.Reorder(projectID, ids)
}
