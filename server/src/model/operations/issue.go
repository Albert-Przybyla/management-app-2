package model_operations

import (
	model_item "menagment-app-2/src/model/item"
	model_package "menagment-app-2/src/model/package"
	"time"
)

type Issue struct {
	ID                string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	SourceStorageID   string    `gorm:"type:uuid;not null" json:"source_storage_id,omitempty"`
	DestinationUserID string    `gorm:"type:uuid;not null" json:"user_id,omitempty"`
	Notes             *string   `gorm:"size:255" json:"notes"`
	OrganizationID    string    `gorm:"type:uuid" json:"organization_id"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Items    []model_item.Item       `gorm:"many2many:issue_items;foreignKey:ID;joinForeignKey:IssueID;References:ID;joinReferences:ItemID" json:"items"`
	Packages []model_package.Package `gorm:"many2many:issue_packages;foreignKey:ID;joinForeignKey:IssueID;References:ID;joinReferences:PackageID" json:"packages"`
}

type IssueItem struct {
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	IssueID string `gorm:"type:uuid;not null" json:"transaction_id"`
	ItemID  string `gorm:"type:uuid;not null" json:"item_id"`
}

type IssuePackage struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	IssueID   string `gorm:"type:uuid;not null" json:"transaction_id"`
	PackageID string `gorm:"type:uuid;not null" json:"package_id"`
}
