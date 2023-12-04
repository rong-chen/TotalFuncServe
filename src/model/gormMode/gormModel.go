package gormMode

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
