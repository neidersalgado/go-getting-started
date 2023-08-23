package models

type Favorite struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserID int    `json:"user_id" gorm:"type:int;not null"`
	RefAPI string `json:"ref_api" gorm:"type:varchar(255);not null"`
}
