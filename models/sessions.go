package models

import (
	"time"
)

type Session struct {
	ID        uint `gorm:"primaryKey"`
	User      User `gorm:"foreignKey:UserID;references:ID"`
	UserID    uint
	StartTime time.Time `gorm:"autoCreateTime"`
	EndTime   *time.Time
}
