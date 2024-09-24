package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
	DeletedAt bool      `json:"-" gorm:"default:false"`
}
