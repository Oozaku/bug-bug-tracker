package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username  string `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex"`
	Name      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Histories []History      `gorm:"foreignKey:UserUsername"`
}
