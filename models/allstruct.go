package models

//type User struct {
//	ID        uint      `json:"id" gorm:"primaryKey"`
//	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
//	UserName  string    `json:"username" gorm:"type:varchar(100);unique;not null"`
//	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
//	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
//	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
//	Role      string    `json:"role" gorm:"type:varchar(255);not null"`
//	BaseModel
//}
//
//type Resume struct {
//	ID              uint        `json:"id"`
//	UserID          uint        `json:"user_id"`
//	FullName        string      `json:"full_name"`
//	Summary         string      `json:"summary"`
//	Skills          string      `json:"skills"`
//	ExperienceYears uint        `json:"experience_years"`
//	Education       string      `json:"education"`
//	Certifications  string      `json:"certifications"`
//	Location        string      `json:"location"`
//	JobCategoryID   uint        `json:"job_category_id"`
//	JobCategory     JobCategory `gorm:"foreignKey:JobCategoryID"`
//	BaseModel
//}
//
//type Job struct {
//	ID            uint        `json:"id" gorm:"primaryKey"`
//	Title         string      `json:"title"`
//	Description   string      `json:"description"`
//	Location      string      `json:"location"`
//	Salary        string      `json:"salary"`
//	CompanyID     uint        `json:"company_id"`
//	Company       Company     `json:"company" gorm:"foreignKey:CompanyID"`
//	JobCategoryID uint        `json:"job_category_id"`
//	JobCategory   JobCategory `json:"job_category" gorm:"foreignKey:JobCategoryID"`
//	BaseModel
//}
//
//type JobCategory struct {
//	ID   uint   `json:"id"`
//	Name string `json:"name"`
//	BaseModel
//}
//
//type Favorite struct {
//	ID     uint `json:"id" gorm:"primaryKey"`
//	UserID uint `json:"user_id" gorm:"not null"`
//	User   User `json:"user" gorm:"foreignKey:UserID"`
//	JobID  uint `json:"job_id" gorm:"not null"`
//	Job    Job  `json:"job" gorm:"foreignKey:JobID"`
//	BaseModel
//}
//
//type Company struct {
//	ID          uint   `json:"id" gorm:"primaryKey"`
//	Name        string `json:"name" gorm:"type:varchar(100);unique;not null"`
//	Description string `json:"description" gorm:"type:text"`
//	BaseModel
//}
//
//type BaseModel struct {
//	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
//	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
//	DeletedAt bool      `json:"deleted_at" gorm:"default:false"`
//}
//
//type Application struct {
//	ID     uint   `json:"id" gorm:"primaryKey"`
//	UserID uint   `json:"user_id" gorm:"not null"`
//	User   User   `json:"user" gorm:"foreignKey:UserID"`
//	JobID  uint   `json:"job_id" gorm:"not null"`
//	Job    Job    `json:"job" gorm:"foreignKey:JobID"`
//	Resume string `json:"resume" gorm:"not null"`
//	Status string `json:"status" gorm:"not null"`
//	BaseModel
//}
