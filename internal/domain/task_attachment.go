package domain

type TaskAttachment struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	TaskID   int64  `gorm:"not null" json:"taskId"`
	FileURL  string `gorm:"type:text;not null" json:"fileUrl"`
	FileName string `json:"fileName"`
	AppModel
}
