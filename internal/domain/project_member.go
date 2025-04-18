package domain

type ProjectMember struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	ProjectID int64  `gorm:"not null" json:"projectId"`
	UserID    int64  `gorm:"not null" json:"userId"`
	Role      string `gorm:"size:50;not null" json:"role"`
	AppModel
}
