package model_storage

type Storage struct {
	Id             string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name           string `gorm:"not null;size:255" json:"name"`
	OrganizationID string `gorm:"type:uuid" json:"organization_id"`
}
