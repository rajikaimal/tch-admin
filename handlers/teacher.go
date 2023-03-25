package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
)

type TeacherHandler struct{}

type SuspendRequestBody struct {
	Email string
}

func (t TeacherHandler) RegisterStudents(c *gin.Context) {
	student := models.Student{Id: 2, Name: "John"}
	db.DB.Create(&student)

	c.IndentedJSON(http.StatusOK, nil)
}

func (t TeacherHandler) SuspendStudent(c *gin.Context) {
	var requestBody SuspendRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid JSON"})
		return
	}

	db.DB.Model(&models.Student{}).
		Where("email = ?", requestBody.Email).
		Update("suspended", true)

	c.IndentedJSON(http.StatusOK, true)
}
