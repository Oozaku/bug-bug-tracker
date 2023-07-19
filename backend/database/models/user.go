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
	Histories []History
}

type UserPublic struct {
	Username  string
	Name      string
	Role      string
	Histories []History
}

type UserPrivate struct {
	Username  string
	Email     string
	Name      string
	Role      string
	Histories []History
}
