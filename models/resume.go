package models

type Resume struct {
	ID              uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          uint        `json:"user_id" gorm:"not null"`
	FullName        string      `json:"full_name" gorm:"type:varchar(255);not null"`
	Summary         string      `json:"summary" gorm:"type:text"`
	Skills          string      `json:"skills" gorm:"type:text"`
	ExperienceYears uint        `json:"experience_years"`
	Education       string      `json:"education" gorm:"type:text"`
	Certifications  string      `json:"certifications" gorm:"type:text"`
	Location        string      `json:"location" gorm:"type:varchar(255)"`
	JobCategoryID   uint        `json:"job_category_id" gorm:"not null"`
	JobCategory     JobCategory `gorm:"foreignKey:JobCategoryID"`
	BaseModel
}
