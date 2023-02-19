package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Email       string `gorm:"primaryKey"`
	IsSuspended bool
	Teachers    []*Teacher `gorm:"many2many:teacher_students;"`
}
