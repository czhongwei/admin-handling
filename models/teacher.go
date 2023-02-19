package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	ID    int
	Email string `gorm:"embedded"`
}
