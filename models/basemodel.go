package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt bool      `json:"deleted_at" gorm:"default:false"`
}
