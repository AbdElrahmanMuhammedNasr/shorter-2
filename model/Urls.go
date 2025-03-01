package model

import (
	"gorm.io/gorm"
	"time"
)

type Urls struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	ShortCode   string    `gorm:"uniqueIndex"`
	OriginalURL string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UserID      uint
}
