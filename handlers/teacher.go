package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
)

type TeacherHandler struct{}

type RegisterReqBody struct {
	Teacher  string
	Students []string
}

type SuspendReqBody struct {
	Email string
}

func (t TeacherHandler) RegisterStudents(c *gin.Context) {
	var requestBody RegisterReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid JSON"})
		return
	}

	var newStudents []models.Student

	for _, s := range requestBody.Students {
		newStudents = append(newStudents, models.Student{Email: s})
	}

	var tch models.Teacher
	var std []models.Student

	db.DB.Model(&models.Teacher{}).Where("email = ?", requestBody.Teacher).Find(&tch)
	db.DB.Model(&models.Student{}).Where("email = ?", "student3@gmail.com").Find(&std)

	db.DB.Model(&tch).Association("Students").Append(&std)

	c.Status(http.StatusNoContent)
}

func (t TeacherHandler) SuspendStudent(c *gin.Context) {
	var requestBody SuspendReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid JSON"})
		return
	}

	db.DB.Model(&models.Student{}).
		Where("email = ?", requestBody.Email).
		Update("suspended", true)

	c.Status(http.StatusNoContent)
}
