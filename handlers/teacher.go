package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/rajikaimal/tch-admin/models"
	respository "github.com/rajikaimal/tch-admin/repository"
	"github.com/rajikaimal/tch-admin/utils"
)

type TeacherHandler struct {
	TeacherRepo respository.TRepo
}

type CommonStudentReqQuery struct {
	Teacher []string
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

// POST /api/register
// register students to a teacher
func (t TeacherHandler) RegisterStudents(c *gin.Context) {
	var requestBody RegisterReqBody

	// check if request body exists
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// validate request body
	if err := requestBody.ValidateRegisterReqBody(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// save to DB
	if err := t.RegisterStudentsHandler(&requestBody.Students, &requestBody.Teacher); err != nil {
		fmt.Println("Error when registering student: ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// handle business logic in registering a student
func (t TeacherHandler) RegisterStudentsHandler(students *[]string, teacher *string) error {
	var newStudents []models.Student

	for _, s := range *students {
		newStudents = append(newStudents, models.Student{Email: s})
	}

	var tch models.Teacher
	var stds []models.Student

	var fields []string
	var values []interface{}

	for _, e := range *students {
		fields = append(fields, fmt.Sprintf("email = ?"))
		values = append(values, e)
	}

	if err := t.TeacherRepo.FindTeacher(*teacher, &tch); err != nil {
		fmt.Println("Error when finding teacher: ", err)
		return errors.New("Couldn't find teacher")
	}

	if tch.Email == "" {
		return errors.New("Teacher not found")
	}

	if err := t.TeacherRepo.FindStudent(fields, values, &stds); err != nil {
		fmt.Println("Error when finding student: ", err)
		return errors.New("Couldn't find student")
	}

	if len(stds) == 0 {
		return errors.New("Student(s) not found")
	}

	if err := t.TeacherRepo.RegisterStudent(&tch, &stds); err != nil {
		fmt.Println("Error when processing: ", err)
		return errors.New("Error when registering student(s)")
	}

	return nil
}

// GET /api/commonstudents
// common students between teachers
func (t TeacherHandler) GetCommonStudents(c *gin.Context) {
	qParam := CommonStudentReqQuery{Teacher: c.QueryArray("teacher")}

	// validate request query param
	if err := qParam.ValidateCommonStudentReqBody(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	var commonStudents Student
	var students []models.Student

	// get common students from DB
	if err := t.GetCommonStudentHandler(qParam.Teacher, &commonStudents, students); err != nil {
		fmt.Println("Error when processing", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	if len(commonStudents.Students) == 0 {
		c.IndentedJSON(http.StatusOK, Student{Students: []string{}})
		return
	}

	c.IndentedJSON(http.StatusOK, commonStudents)
}

// handle business logic in getting common students
func (t TeacherHandler) GetCommonStudentHandler(teachers []string, commonStudents *Student, students []models.Student) error {
	var fields []string
	var values []interface{}

	for _, e := range teachers {
		fields = append(fields, fmt.Sprintf("teachers.email = ?"))
		values = append(values, e)
	}

	if err := t.TeacherRepo.GetCommonStudents(fields, values, &students); err != nil {
		return errors.New("Error getting asd common students")
	}

	for _, s := range students {
		commonStudents.Students = append(commonStudents.Students, s.Email)
	}

	return nil
}

// POST /api/suspend
// suspend a student
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

	// update suspended field of student in DB
	if err := t.SuspendStudentHandler(requestBody.Email); err != nil {
		fmt.Println("Error when processing", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// handle business logic in suspending a student
func (t TeacherHandler) SuspendStudentHandler(email string) error {
	var student models.Student
	// check if student exists
	if err := t.TeacherRepo.FindOneStudent(email, &student); err != nil {
		fmt.Println("Error when finding student: ", err)
		return errors.New("Couldn't find student")
	}

	if student.Email == "" {
		return errors.New("Student is not registered with the system")
	}

	if student.Suspended {
		return errors.New("Student is already suspended")
	}

	// update sesponded field of student
	if err := t.TeacherRepo.SuspendStudent(email); err != nil {
		fmt.Println("Error when processing", err)
		return errors.New("Error when suspending student")
	}

	return nil
}

// POST /api/retrievefornotifications
// retrieve students who can receive a given notification
func (t TeacherHandler) RetrieveNotifications(c *gin.Context) {
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

	var allRecipients []string
	if err := t.RetrieveNotificationsHandler(teacher, notificationTxt, &allRecipients); err != nil {
		fmt.Println("Error when finding students: ", err)
		c.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, Recipient{Recipients: allRecipients})
}

// handle business logic in retrieving students who can receive a given notification
func (t TeacherHandler) RetrieveNotificationsHandler(teacher string, notificationTxt string, allRecipients *[]string) error {
	var students []models.Student
	if err := t.TeacherRepo.RetrieveNotifications(teacher, &students); err != nil {
		fmt.Println("Error when finding students: ", err)
		return errors.New("Error when finding students")
	}

	for _, s := range students {
		*allRecipients = append(*allRecipients, s.Email)
	}

	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)
	mentions := re.FindAllString(notificationTxt, -1)

	var fields []string
	var values []interface{}

	var studentsToRecieve []models.Student

	for _, email := range mentions {
		fields = append(fields, fmt.Sprintf("email = ?"))
		values = append(values, email)
	}

	// check all the mentioned emails in Notification text exist
	if err := t.TeacherRepo.FindStudentsForNotifications(fields, values, false, &studentsToRecieve); err != nil {
		fmt.Println("Error when finding student: ", err)
		return errors.New("Couldn't find student")
	}

	for _, s := range studentsToRecieve {
		if !utils.Contains(*allRecipients, s.Email) {
			*allRecipients = append(*allRecipients, s.Email)
		}
	}

	return nil
}
