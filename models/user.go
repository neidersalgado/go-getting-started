package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"` // El "-" evita que se serialice al enviarlo en JSON
	Address   string    `json:"address" gorm:"type:varchar(255);not null"`
	Birthdate time.Time `json:"birthdate" gorm:"type:date;not null"`
	City      string    `json:"city" gorm:"type:varchar(255);not null"`
}
