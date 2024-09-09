package models

type Application struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id" gorm:"not null"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	JobID    uint   `json:"job_id" gorm:"not null"`
	Job      Job    `json:"job" gorm:"foreignKey:JobID"`
	ResumeID uint   `json:"resume_id" gorm:"not null"`
	Resume   Resume `json:"resume" gorm:"foreignKey:ResumeID"`
	Status   string `json:"status" gorm:"not null;default:'pending'"`
	BaseModel
}
