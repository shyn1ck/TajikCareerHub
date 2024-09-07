package models

type JobCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(100);unique;not null"`
	BaseModel
}
