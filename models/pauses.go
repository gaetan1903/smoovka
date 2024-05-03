package models

import (
	"time"
)

type Pause struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Session   Session `gorm:"foreignKey:SessionID;references:ID"`
	SessionID uint
	StartTime time.Time `gorm:"autoCreateTime"`
	EndTime   *time.Time
}
