package models

type Application struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	JobID  uint   `json:"job_id" gorm:"not null"`
	Job    Job    `json:"job" gorm:"foreignKey:JobID"`
	Resume string `json:"resume" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
	BaseModel
}
