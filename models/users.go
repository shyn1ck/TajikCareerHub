package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
	UserName  string    `json:"username" gorm:"type:varchar(100);unique;not null"`
	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"type:varchar(255);not null"`
	BaseModel
}
