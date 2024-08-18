package models

import (
	"time"
)

type Company struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"varchar(100);unique"`
	Description string    `json:"description" gorm:"text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   bool      `json:"deleted_at" gorm:"default:false"`
}
