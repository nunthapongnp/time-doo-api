package model

import "time-doo-api/internal/domain"

type UserDTO struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	FullName string `json:"fullName,omitempty"`
	TenantID int64  `json:"tenantId,omitempty"`
	Role     string `json:"role,omitempty"`
	domain.AppModel
}
