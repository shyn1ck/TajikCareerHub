package models

type Application struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id" gorm:"not null"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	VacancyID uint    `json:"vacancy_id" gorm:"not null"`
	Vacancy   Vacancy `json:"vacancy" gorm:"foreignKey:VacancyID"`
	ResumeID  uint    `json:"resume_id" gorm:"not null"`
	Resume    Resume  `json:"resume" gorm:"foreignKey:ResumeID"`
	Status    string  `json:"status" gorm:"not null;default:'pending'"`
	BaseModel
}

type ApplicationStatusUpdate struct {
	ApplicationID uint   `json:"application_id" binding:"required"`
	Status        string `json:"status" binding:"required"`
}
