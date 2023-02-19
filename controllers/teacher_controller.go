package controllers

import (
	"admin-handling/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherEmail struct {
	Teacher string `json:"teacher"`
}

type Register struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}

func AddTeacher(c *gin.Context) {
	var teacher TeacherEmail
	if err := c.BindJSON(&teacher); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
		return
	}
	newTeacher := models.Teacher{Email: teacher.Teacher}

	result := db.Create(&newTeacher)
	if err := result.Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to add teacher"})
		return
	}
	c.Status(http.StatusCreated)
}

func RegisterStudents(c *gin.Context) {
	var request Register
	var teacher models.Teacher
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
		return
	}
	result := db.Where("email = ?", request.Teacher).First(&teacher)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "teacher does not exist"})
		return
	}
	for _, student := range request.Students {
		var studentToRegister models.Student
		result := db.Where("email = ?", student).First(&studentToRegister)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "student does not exist"})
			return
		}
		db.Model(&teacher).Association("Students").Append([]models.Student{studentToRegister})
	}

	c.Status(http.StatusNoContent)
}
