package domain

import (
	"time"
	"time-doo-api/internal/middleware"

	"gorm.io/gorm"
)

type AppModel struct {
	CreatedBy int64      `json:"createdBy"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedBy *int64     `gorm:"default:null" json:"updatedBy"`
	UpdatedAt *time.Time `gorm:"default:null" json:"updatedAt"`
}

func (m *AppModel) BeforeCreate(db *gorm.DB) (err error) {
	userID, ok := db.Statement.Context.Value(middleware.UserCtxKey).(int64)
	if ok {
		m.CreatedBy = userID
	}
	m.CreatedAt = time.Now()
	return
}

func (m *AppModel) BeforeUpdate(db *gorm.DB) (err error) {
	userID, ok := db.Statement.Context.Value(middleware.UserCtxKey).(int64)
	if ok {
		m.UpdatedBy = &userID
	}
	now := time.Now()
	m.UpdatedAt = &now
	return
}
