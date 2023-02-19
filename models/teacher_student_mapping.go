package models

import "gorm.io/gorm"

type TeacherStudentMapping struct {
	gorm.Model
	Teacher string `gorm:"primaryKey"`
	Student string `gorm:"primaryKey"`
}
