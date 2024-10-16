package model_item

import "time"

type Item struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name           string    `gorm:"not null;size:255" json:"name"`
	Description    *string   `gorm:"size:255" json:"description"`
	Code           *string   `gorm:"size:255" json:"code"`
	Quantity       int       `gorm:"not null" json:"quantity"`
	OrganizationID string    `gorm:"type:uuid" json:"organization_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	PackageID      *string   `gorm:"type:uuid" json:"package_id,omitempty"`
}
