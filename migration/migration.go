package main

import (
	"admin-handling/database"
	"admin-handling/models"
)

func main() {
	db := database.InitDb()
	db.AutoMigrate(&models.Teacher{})
	db.AutoMigrate(&models.Student{})
}
