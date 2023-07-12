package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	ActionDescription string
	ChangedTo         string
	IssueID           uint
	UserUsername      string
}
