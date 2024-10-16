package model_organization

import (
	model_storage "menagment-app-2/src/model/storage"
	model_user "menagment-app-2/src/model/user"
	"time"
)

type Organization struct {
	ID        string                  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string                  `gorm:"unique;not null;size:255" json:"name" binding:"required,min=3,max=255"`
	CreatedAt time.Time               `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time               `gorm:"autoUpdateTime" json:"updated_at"`
	Users     []model_user.User       `gorm:"foreignKey:OrganizationID" json:"users"`
	Storages  []model_storage.Storage `gorm:"foreignKey:OrganizationID" json:"storages"`
}
