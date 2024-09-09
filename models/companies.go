package models

type Company struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(100);unique;not null"`
	Description string `json:"description" gorm:"type:text"`
	BaseModel
}
