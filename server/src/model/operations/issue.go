package model_operations

import "time"

type Issue struct {
	ID                string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	SourceStorageID   string    `gorm:"type:uuid" json:"source_storage_id,omitempty"`
	DestinationUserID string    `gorm:"type:uuid" json:"user_id,omitempty"`
	Notes             *string   `gorm:"size:255" json:"notes"`
	OrganizationID    string    `gorm:"type:uuid" json:"organization_id"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type IssueItem struct {
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	IssueID  string `gorm:"type:uuid;not null" json:"transaction_id"`
	ItemID   string `gorm:"type:uuid;not null" json:"item_id"`
	Quantity int    `gorm:"not null" json:"quantity"`
}

type IssuePackage struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	IssueID   string `gorm:"type:uuid;not null" json:"transaction_id"`
	PackageID string `gorm:"type:uuid;not null" json:"package_id"`
	Quantity  int    `gorm:"not null" json:"quantity"`
}
