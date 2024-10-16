package model_transaction

import (
	"time"
)

type TransactionType string

const (
	TransactionIssue    TransactionType = "issue"
	TransactionReturn   TransactionType = "return"
	TransactionTransfer TransactionType = "transfer"
)

type Transaction struct {
	Id              string          `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	SourceStorageID *string         `gorm:"type:uuid" json:"source_storage_id,omitempty"`
	UserID          *string         `gorm:"type:uuid" json:"user_id,omitempty"`
	TransactionType TransactionType `gorm:"type:varchar(50);not null" json:"transaction_type"`
	Notes           *string         `gorm:"size:255" json:"notes"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

type TransactionItem struct {
	Id            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TransactionID string `gorm:"type:uuid;not null" json:"transaction_id"`
	ItemID        string `gorm:"type:uuid;not null" json:"item_id"`
	Quantity      int    `gorm:"not null" json:"quantity"`
}

type TransactionPackage struct {
	Id            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TransactionID string `gorm:"type:uuid;not null" json:"transaction_id"`
	PackageID     string `gorm:"type:uuid;not null" json:"package_id"`
	Quantity      int    `gorm:"not null" json:"quantity"`
}
