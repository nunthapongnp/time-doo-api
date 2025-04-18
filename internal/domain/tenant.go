package domain

type Tenant struct {
	ID   int64  `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	AppModel
}
