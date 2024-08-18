package models

import "time"

type Job struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Location    string    `json:"location" gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   bool      `json:"deleted_at" gorm:"default:false"`
	Company
	JobCategory
}
