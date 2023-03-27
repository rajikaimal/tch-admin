package handlers

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rajikaimal/tch-admin/mocks"
	"github.com/rajikaimal/tch-admin/models"
)

func TestRegisterStudentsHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockTeacherRepo := mocks.NewMockTRepo(mockCtrl)

	var teacher models.Teacher
	var students []models.Student

	teacherEmail := "teacher1@gmail.com"

	mockTeacherRepo.EXPECT().FindTeacher(teacherEmail, &teacher).AnyTimes()
	mockTeacherRepo.EXPECT().FindStudent([]string{"email = ?"}, []interface{}{"student1@gmail.com"}, &students).Do(func(fields []string, values []interface{}, stds *[]models.Student) {
		*stds = append(*stds, models.Student{Id: 1, Email: "student1@gmail.com"})
	}).AnyTimes()

	mockTeacherRepo.EXPECT().RegisterStudent(&teacher, &students).Return(nil).AnyTimes()

	stdsToRegiser := []string{"student1@gmail.com"}

	th := TeacherHandler{TeacherRepo: mockTeacherRepo}
	err := th.RegisterStudentsHandler(&stdsToRegiser, &teacherEmail)

	if err != nil {
		fmt.Println(err.Error())
		panic("Registering failed")
	}
}

func TestGetCommonStudentsHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockTeacherRepo := mocks.NewMockTRepo(mockCtrl)
	var students []models.Student

	mockTeacherRepo.EXPECT().GetCommonStudents([]string{"teachers.email = ?"}, []interface{}{"teacher1@gmail.com"}, &students).Do(func(fields []string, values []interface{}, stds *[]models.Student) {
		*stds = append(*stds, models.Student{Id: 1, Email: "student1@gmail.com"})
	})

	var commonStudents Student
	want := models.Student{Id: 1, Email: "student1@gmail.com"}

	th := TeacherHandler{TeacherRepo: mockTeacherRepo}

	th.GetCommonStudentHandler([]string{"teacher1@gmail.com"}, &commonStudents, students)

	fmt.Println(commonStudents)

	if commonStudents.Students[0] != want.Email {
		panic("failed common students handler")
	}
}

func TestSuspendStudentHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockTeacherRepo := mocks.NewMockTRepo(mockCtrl)
	mockTeacherRepo.EXPECT().SuspendStudent("student1@gmail.com").Return(nil)

	var student models.Student

	mockTeacherRepo.EXPECT().FindOneStudent("student1@gmail.com", &student).Do(func(email string, student *models.Student) {
		student.Id = 1
		student.Email = "student1@gmail.com"
	})

	th := TeacherHandler{TeacherRepo: mockTeacherRepo}

	err := th.SuspendStudentHandler("student1@gmail.com")

	if err != nil {
		panic("Failed to suspend student")
	}
}

func TestRetrieveNotificationsHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockTeacherRepo := mocks.NewMockTRepo(mockCtrl)
	var students []models.Student

	mockTeacherRepo.EXPECT().RetrieveNotifications("teacher1@gmail.com", &students).Do(func(fields string, stds *[]models.Student) {
		*stds = append(*stds, models.Student{Id: 1, Email: "student1@gmail.com"})
	})

	var studentsToRecieve []models.Student

	mockTeacherRepo.EXPECT().FindStudentsForNotifications([]string{"email = ?"}, []interface{}{"student1@gmail.com"}, false, &studentsToRecieve).Do(func(fields []string, values []interface{}, uspended bool, students *[]models.Student) {
		*students = append(*students, models.Student{Id: 1, Email: "student1@gmail.com"})
	})

	th := TeacherHandler{TeacherRepo: mockTeacherRepo}

	var allRecipients []string

	th.RetrieveNotificationsHandler("teacher1@gmail.com", "Hello @student1@gmail.com", &allRecipients)

	fmt.Println(allRecipients)
	want := []string{"student1@gmail.com", "student2@gmail.com"}

	if want[0] != allRecipients[0] {
		panic("didn't match student in notifications")
	}
}
