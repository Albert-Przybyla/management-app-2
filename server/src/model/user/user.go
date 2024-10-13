package model_user

import "time"

type User struct {
	Id           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email        string    `gorm:"unique;not null;size:255"`
	Password     string    `gorm:"not null;size:255"`
	FirstName    string    `gorm:"size:100"`
	LastName     string    `gorm:"size:100"`
	RefreshToken string    `gorm:"size:255"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
