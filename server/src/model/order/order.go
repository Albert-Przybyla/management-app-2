package model_order

import "time"

type Order struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Quantity       int       `gorm:"not null" json:"quantity"`
	OrganizationID string    `gorm:"type:uuid" json:"organization_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
