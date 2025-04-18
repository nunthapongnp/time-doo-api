package domain

type Project struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	TenantID    int64  `gorm:"not null" json:"tenantId"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	AppModel
}
