package model_package

import (
	model_item "menagment-app-2/src/model/item"
	"time"
)

type Package struct {
	Id        string            `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string            `gorm:"not null;size:255" json:"name"`
	ItemID    string            `gorm:"type:uuid;not null" json:"item_id"`
	Items     []model_item.Item `gorm:"foreignKey:PackageID" json:"items"`
	CreatedAt time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
}
