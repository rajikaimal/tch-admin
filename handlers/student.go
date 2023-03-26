package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
	"github.com/rajikaimal/tch-admin/utils"
)

type StudentHandler struct{}

type Student struct {
	Students []string `json:"students"`
}

type Recipient struct {
	Recipients []string `json:"recipients"`
}

type RetrieveNotificationReqBody struct {
	Teacher      string
	Notification string
}

func (h StudentHandler) GetCommonStudents(c *gin.Context) {
	var commonStudents Student

	qParam := c.Query("teacher")

	if qParam == "" {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Missing teacher query parameter"})
		return
	}

	emails := []string{qParam}

	var fields []string
	var values []interface{}

	for _, e := range emails {
		fields = append(fields, fmt.Sprintf("teachers.email = ?"))
		values = append(values, e)
	}

	var students []models.Student
	db.DB.Model(&models.Student{}).
		Select("DISTINCT students.email").
		Joins("JOIN registers ON registers.student_id = students.id AND registers.student_email = students.email").
		Joins("JOIN teachers ON registers.teacher_id = teachers.id AND registers.student_email = students.email").
		Where(strings.Join(fields, " OR "), values...).
		Find(&students)

	for _, s := range students {
		commonStudents.Students = append(commonStudents.Students, s.Email)
	}

	c.IndentedJSON(http.StatusOK, commonStudents)
}

func (h StudentHandler) RetrieveNotifications(c *gin.Context) {
	var requestBody RetrieveNotificationReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid JSON"})
		return
	}

	teacher := requestBody.Teacher
	notificationTxt := requestBody.Notification

	if teacher == "" || notificationTxt == "" {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Missing teacher or notification text in request"})
		return
	}

	var students []models.Student
	db.DB.Model(&models.Student{}).
		Select("DISTINCT students.email").
		Joins("JOIN registers ON registers.student_id = students.id AND registers.student_email = students.email").
		Joins("JOIN teachers ON registers.teacher_id = teachers.id AND registers.student_email = students.email").
		Where("registers.teacher_email = ?", teacher).
		Find(&students)

	var allRecipients []string

	for _, s := range students {
		allRecipients = append(allRecipients, s.Email)
	}

	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)
	mentions := re.FindAllString(notificationTxt, -1)

	for _, m := range mentions {
		if !utils.Contains(allRecipients, m) {
			allRecipients = append(allRecipients, m)
		}
	}

	c.IndentedJSON(http.StatusOK, Recipient{Recipients: allRecipients})
}
