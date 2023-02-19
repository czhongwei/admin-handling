package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID           int
	Email        string `gorm:"unique"`
	IsSuspended  bool
	TeacherRefer int
	Teacher      Teacher `gorm:"foreignKey:TeacherRefer"`
}
