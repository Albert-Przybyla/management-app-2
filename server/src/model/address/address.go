package model_address

import "time"

type Customer struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name           string    `gorm:"not null;size:255" json:"name"`
	OrganizationID string    `gorm:"type:uuid" json:"organization_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
