package domain

type TenantMember struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	TenantID int64  `gorm:"not null" json:"tenantId"`
	UserID   int64  `gorm:"not null" json:"userId"`
	Role     string `gorm:"size:50;not null" json:"role"`
	AppModel
	Users   User   `gorm:"foreignKey:UserID" json:"user"`
	Tenants Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
}
