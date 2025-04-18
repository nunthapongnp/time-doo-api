package domain

import (
	"gorm.io/datatypes"
)

type ActivityLog struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	ProjectID int64          `gorm:"not null" json:"projectId"`
	UserID    int64          `json:"userId"`
	Action    string         `gorm:"size:255" json:"action"`
	Metadata  datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
	AppModel
}
