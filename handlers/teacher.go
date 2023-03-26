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

type TeacherHandler struct{}

type CommonStudentReqQuery struct {
	Teacher string
}

type RegisterReqBody struct {
	Teacher  string
	Students []string
}

type SuspendReqBody struct {
	Email string
}

type RetrieveNotificationReqBody struct {
	Teacher      string
	Notification string
}

type Student struct {
	Students []string `json:"students"`
}

type Recipient struct {
	Recipients []string `json:"recipients"`
}

func (t TeacherHandler) GetCommonStudents(c *gin.Context) {
	qParam := CommonStudentReqQuery{Teacher: c.Query("teacher")}

	// check if request query param exists
	if qParam.Teacher == "" {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: "Missing teacher query parameter"})
		return
	}

	// validate request query param
	if err := qParam.ValidateCommonStudentReqBody(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	emails := []string{qParam.Teacher}

	var fields []string
	var values []interface{}

	for _, e := range emails {
		fields = append(fields, fmt.Sprintf("teachers.email = ?"))
		values = append(values, e)
	}

	var commonStudents Student
	var students []models.Student

	if err := db.DB.Model(&models.Student{}).
		Select("DISTINCT students.email").
		Joins("JOIN registers ON registers.student_id = students.id AND registers.student_email = students.email").
		Joins("JOIN teachers ON registers.teacher_id = teachers.id AND registers.student_email = students.email").
		Where(strings.Join(fields, " OR "), values...).
		Find(&students).Error; err != nil {
		fmt.Println("Error when finding student: ", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: "Error when finding students"})
	}

	for _, s := range students {
		commonStudents.Students = append(commonStudents.Students, s.Email)
	}

	c.IndentedJSON(http.StatusOK, commonStudents)
}

func (t TeacherHandler) RegisterStudents(c *gin.Context) {
	var requestBody RegisterReqBody

	// check if request body exists
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: InvalidRequBodyErr})
		return
	}

	// validate request body
	if err := requestBody.ValidateRegisterReqBody(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	var newStudents []models.Student

	for _, s := range requestBody.Students {
		newStudents = append(newStudents, models.Student{Email: s})
	}

	var tch models.Teacher
	var stds []models.Student

	var fields []string
	var values []interface{}

	for _, e := range requestBody.Students {
		fields = append(fields, fmt.Sprintf("email = ?"))
		values = append(values, e)
	}

	if err := db.DB.Model(&models.Teacher{}).Where("email = ?", requestBody.Teacher).Find(&tch).Error; err != nil {
		fmt.Println("Error when finding teacher: ", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: "Couldn't find teacher"})
		return
	}

	if err := db.DB.Model(&models.Student{}).Where(strings.Join(fields, " OR "), values...).Find(&stds).Error; err != nil {
		fmt.Println("Error when finding student: ", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: "Couldn't find student"})
		return
	}

	db.DB.Model(&tch).Association("Students").Append(&stds)

	c.Status(http.StatusNoContent)
}

func (t TeacherHandler) SuspendStudent(c *gin.Context) {
	var requestBody SuspendReqBody

	// check if request body exists
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: InvalidRequBodyErr})
		return
	}

	// validate request body
	if err := requestBody.ValidateSuspendReqBody(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// update sesponded field of student
	if err := db.DB.Model(&models.Student{}).
		Where("email = ?", requestBody.Email).
		Update("suspended", true).Error; err != nil {
		fmt.Println("Error when processing", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: "Error when suspeding student"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h TeacherHandler) RetrieveNotifications(c *gin.Context) {
	var requestBody RetrieveNotificationReqBody

	// check if request body exists
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: InvalidRequBodyErr})
		return
	}

	// validate request body
	if err := requestBody.ValidateRetrieveNotificationReqBody(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
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
