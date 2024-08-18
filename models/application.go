package models

import (
	"gorm.io/gorm"
)

type Application struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	JobID     uint           `json:"job_id" gorm:"not null"`
	Resume    string         `json:"resume" gorm:"not null"`
	Status    string         `json:"status" gorm:"not null"`
	CreatedAt gorm.DeletedAt `json:"created_at" gorm:"index"`
	UpdatedAt gorm.DeletedAt `json:"updated_at" gorm:"index"`
	DeletedAt bool           `json:"deleted_at" gorm:"default:false"`
}
