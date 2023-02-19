package controllers

import (
	"errors"
	"net/http"

	"admin-handling/database"
	"admin-handling/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RetrieveNotificationsArgs struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}

type StudentEmail struct {
	Student string `json:"student"`
}

var db *gorm.DB = database.InitDb()

func SuspendStudent(c *gin.Context) {
	var student StudentEmail
	var studentToSuspend models.Student
	// Call BindJSON to bind the received JSON to student
	if err := c.BindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
	}

	// search for student in DB set IsSuspended to true
	result := db.First(&studentToSuspend, "email = ?", student.Student)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "student does not exist"})
	}
	studentToSuspend.IsSuspended = true
	db.Save(&studentToSuspend)
	c.Status(http.StatusNoContent)
}

func RetrieveNotifications(c *gin.Context) {
	// TODO
	var request RetrieveNotificationsArgs
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
	}

	// search for teacher in DB

	// comb through notification to search for tagged students

	c.IndentedJSON(http.StatusOK, request)
}

func RegisterStudents(c *gin.Context) {
	// TODO
	//var request RegisterStudents
	c.Status(http.StatusNoContent)
}

func GetCommonStudents(c *gin.Context) {
	// TODO
	students := 1 // placeholder

	c.IndentedJSON(http.StatusOK, students)
}
