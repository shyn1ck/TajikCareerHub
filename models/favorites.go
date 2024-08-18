package models

import (
	"time"
)

type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	JobID     uint      `json:"job_id" gorm:"not null" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt bool      `json:"deleted_at" gorm:"default:false" `
}
