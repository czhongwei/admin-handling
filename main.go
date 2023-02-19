package main

import (
	"admin-handling/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/api/commonstudents", controllers.GetCommonStudents)
	router.POST("/api/teacher", controllers.AddTeacher)
	router.POST("/api/student", controllers.AddStudent)
	router.POST("/api/register", controllers.RegisterStudents)
	router.POST("/api/suspend", controllers.SuspendStudent)
	router.POST("/api/retrievefornotifications", controllers.RetrieveNotifications)

	router.Run("localhost:8080")
}
