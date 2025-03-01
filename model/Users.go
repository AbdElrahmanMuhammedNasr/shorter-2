package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Urls []Urls `gorm:"foreignKey:UserID"`
}
