package models

import (
	"time"
)

type JobCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"varchar(100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt bool      `json:"deleted_at" gorm:"default:false"`
}
