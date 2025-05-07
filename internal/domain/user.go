package domain

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	TenantID  int64  `gorm:"not null" json:"tenantId"`
	Email     string `gorm:"size:255;unique;not null" json:"email"`
	Password  string `gorm:"type:text" json:"-"`
	FullName  string `gorm:"size:255" json:"fullName"`
	AvatarURL string `gorm:"type:text;default:null" json:"avatarUrl"`
	AppModel
}
