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

type Students struct {
	Students []string `json:"students"`
}

var db *gorm.DB = database.InitDb()

func SuspendStudent(c *gin.Context) {
	var student StudentEmail
	var studentToSuspend models.Student
	// Call BindJSON to bind the received JSON to student
	if err := c.BindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
		return
	}

	// search for student in DB set IsSuspended to true
	result := db.First(&studentToSuspend, "email = ?", student.Student)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "student does not exist"})
		return
	}
	studentToSuspend.IsSuspended = true
	db.Save(&studentToSuspend)
	c.Status(http.StatusNoContent)
}

func AddStudent(c *gin.Context) {
	var student StudentEmail
	if err := c.BindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
	}
	newStudent := models.Student{Email: student.Student, IsSuspended: false}

	result := db.Create(&newStudent)
	if err := result.Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to add student"})
		return
	}
	c.Status(http.StatusCreated)
}

func RetrieveNotifications(c *gin.Context) {
	// TODO
	var request RetrieveNotificationsArgs
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid json fields"})
		return
	}

	// search for students under this teacher in DB with is_suspended = false

	// comb through notification to search for tagged students, check is_suspended = false

	c.IndentedJSON(http.StatusOK, request)
}

func GetCommonStudents(c *gin.Context) {
	teachers := c.QueryArray("teacher")
	count := len(teachers)
	if len(teachers) == 0 {
		students := []string{}
		c.IndentedJSON(http.StatusOK, gin.H{"students": students})
		return
	}
	var studentsEmail []string
	db.Table("teacher_students").Select("student_email").Where("teacher_email IN ?", teachers).Group("student_email").Having("COUNT(student_email) = ?", count).Find(&studentsEmail)
	//db.Raw("SELECT TS.student_email FROM teacher_students TS WHERE TS.student_email IN ?", teachers).Group("teacher_students.student_email").Having("COUNT(teacher_students.student_email) = ?", count)

	c.IndentedJSON(http.StatusOK, gin.H{"students": studentsEmail})
}
