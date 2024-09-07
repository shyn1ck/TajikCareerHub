package models

//
//import "time"
//
//type User struct {
//	ID        uint      `json:"id" gorm:"primaryKey"`
//	FullName  string    `json:"full_name" gorm:"type:varchar(255);not null"`
//	UserName  string    `json:"user_name" gorm:"type:varchar(100);unique;not null"`
//	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
//	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
//	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
//	Role      string    `json:"role" gorm:"type:varchar(255);not null"`
//	BaseModel
//}
//
//type Job struct {
//	ID            uint        `json:"id" gorm:"primaryKey"`
//	Title         string      `json:"title" gorm:"type:varchar(100);not null"`
//	Description   string      `json:"description" gorm:"type:text;not null"`
//	Location      string      `json:"location" gorm:"type:varchar(100);not null"`
//	CompanyID     uint        `json:"company_id" gorm:"not null"`
//	Company       Company     `json:"company" gorm:"foreignKey:CompanyID"`
//	JobCategoryID uint        `json:"job_category_id" gorm:"not null"`
//	JobCategory   JobCategory `json:"job_category" gorm:"foreignKey:JobCategoryID"`
//	Salary        string      `json:"salary" gorm:"type:varchar(50)"`
//	BaseModel
//}
//
//
//type JobCategory struct {
//	ID   uint   `json:"id" gorm:"primaryKey"`
//	Name string `json:"name" gorm:"varchar(100)"`
//	BaseModel
//}
//
//type Favorite struct {
//	ID     uint `json:"id" gorm:"primaryKey"`
//	UserID uint `json:"user_id" gorm:"not null"`
//	JobID  uint `json:"job_id" gorm:"not null"`
//	BaseModel
//}
//
//type Company struct {
//	ID          uint   `json:"id" gorm:"primaryKey"`
//	Name        string `json:"name" gorm:"varchar(100);unique"`
//	Description string `json:"description" gorm:"text"`
//	BaseModel
//}
//type BaseModel struct {
//	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
//	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
//	DeletedAt bool      `json:"deleted_at" gorm:"default:false"`
//}
//
//type Application struct {
//	ID     uint   `json:"id" gorm:"primaryKey"`
//	UserID uint   `json:"user_id" gorm:"not null"`
//	JobID  uint   `json:"job_id" gorm:"not null"`
//	Resume string `json:"resume" gorm:"not null"`
//	Status string `json:"status" gorm:"not null"`
//	BaseModel
//}
//
