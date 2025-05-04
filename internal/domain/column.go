package domain

type Column struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	ProjectID int64  `gorm:"not null" json:"projectId"`
	Name      string `gorm:"size:100;not null" json:"name"`
	Position  int64  `gorm:"not null" json:"position"`
	Tasks     []Task `gorm:"foreignKey:ColumnID" json:"tasks"`
	AppModel
}
