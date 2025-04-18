package domain

import "time"

type Task struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	ProjectID   int64      `gorm:"not null" json:"projectId"`
	ColumnID    int64      `gorm:"not null" json:"columnId"`
	Title       string     `gorm:"size:255;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	AssigneeID  *int64     `json:"assigneeId"`
	DueDate     *time.Time `json:"dueDate"`
	Position    int        `gorm:"not null" json:"position"`
	AppModel
}
