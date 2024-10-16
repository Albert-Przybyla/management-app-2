package model_order

import "time"

type Order struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
