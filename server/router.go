package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rajikaimal/tch-admin/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	studentHandler := new(handlers.StudentHandler)
	teacherHandler := new(handlers.TeacherHandler)

	v1 := router.Group("/api/")
	{
		v1.POST("/register", teacherHandler.RegisterStudents)
		v1.GET("/commonstudents", studentHandler.GetCommonStudents)
		v1.POST("/suspend", teacherHandler.SuspendStudent)
		v1.POST("/retrievefornotifications", studentHandler.RetrieveNotifications)
	}

	return router
}
