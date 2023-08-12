package models

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	Title          string
	Priority       string
	Description    string
	Progression    string
	Author         User
	AuthorUsername string
	Assignees      []User `gorm:"many2many:issue_assignees;"`
}
