package controllers

import (
	"net/http"

	"admin-handling/models"

	"github.com/gin-gonic/gin"
)

type TeacherEmail struct {
	Teacher string `json:"teacher"`
}

func AddTeacher(c *gin.Context) {
	var teacher TeacherEmail
	if err := c.BindJSON(&teacher); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
	}
	newTeacher := models.Teacher{Email: teacher.Teacher}

	result := db.Create(&newTeacher)
	if err := result.Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to add teacher"})
		return
	}
	c.Status(http.StatusCreated)
}
