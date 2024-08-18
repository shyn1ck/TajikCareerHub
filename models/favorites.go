package models

type Favorite struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`
	JobID  uint `json:"job_id" gorm:"not null"`
	BaseModel
}
