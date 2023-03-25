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

	v1 := router.Group("/api/")
	{
		v1.POST("/register", studentHandler.RegisterStudents)
		v1.GET("/commonstudents", studentHandler.GetCommonStudents)
		v1.GET("/suspend", studentHandler.GetCommonStudents)
		v1.POST("/retrievefornotifications", studentHandler.RetrieveNotifications)
	}

	return router
}
