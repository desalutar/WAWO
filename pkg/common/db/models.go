package db

import "time"

type User struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username     string    `gorm:"uniqueIndex;not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	IsActive     bool      `gorm:"default:true"`
	IsVerified   bool      `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Session struct {
	ID              string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          string    `gorm:"type:uuid;not null;index"`
	RefreshTokenHash string   `gorm:"uniqueIndex;not null"`
	ExpiresAt       time.Time `gorm:"not null"`
	CreatedAt       time.Time
}
