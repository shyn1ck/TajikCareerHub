package models

type Job struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	Title         string      `json:"title" gorm:"type:varchar(100);not null"`
	Description   string      `json:"description" gorm:"type:text;not null"`
	Location      string      `json:"location" gorm:"type:varchar(100);not null"`
	CompanyID     uint        `json:"company_id" gorm:"not null"`
	Company       Company     `json:"company" gorm:"foreignKey:CompanyID"`
	JobCategoryID uint        `json:"job_category_id" gorm:"not null"`
	JobCategory   JobCategory `json:"job_category" gorm:"foreignKey:JobCategoryID"`
	Salary        string      `json:"salary" gorm:"type:varchar(50)"`
	BaseModel
}
