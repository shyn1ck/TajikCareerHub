package models

type Application struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id" gorm:"not null"`
	JobID  uint   `json:"job_id" gorm:"not null"`
	Resume string `json:"resume" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
	BaseModel
}
