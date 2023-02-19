package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/commonstudents", getCommonStudents)
	router.POST("/api/register", registerStudents)
	router.POST("/api/suspend", suspendStudent)
	router.POST("/api/retrievefornotifications", retrieveNotifications)

	router.Run("localhost:8080")
}

func getCommonStudents(c *gin.Context) {
	// TODO
	students := 1 // placeholder

	c.IndentedJSON(http.StatusOK, students)
}

func registerStudents(c *gin.Context) {
	// TODO
	c.Status(http.StatusNoContent)
}

func suspendStudent(c *gin.Context) {
	// TODO
	c.Status(http.StatusNoContent)
}

func retrieveNotifications(c *gin.Context) {
	// TODO
	students := 1
	c.IndentedJSON(http.StatusOK, students)
}
