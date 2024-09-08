package models

type Job struct {
	ID            uint        `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	Location      string      `json:"location"`
	Salary        int         `json:"salary"`
	CompanyID     uint        `json:"company_id"`
	Company       Company     `json:"company" gorm:"foreignKey:CompanyID"`
	JobCategoryID uint        `json:"job_category_id"`
	JobCategory   JobCategory `json:"job_category" gorm:"foreignKey:JobCategoryID"`
	BaseModel
}
