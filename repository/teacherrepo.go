package respository

import (
	"errors"
	"strings"

	"github.com/rajikaimal/tch-admin/models"
	"gorm.io/gorm"
)

type TRepo interface {
	FindTeacher(email string, tch *models.Teacher) (err error)
	FindStudent(fields []string, values []interface{}, students *[]models.Student) (err error)
	FindOneStudent(email string, student *models.Student) (err error)
	FindStudentsForNotifications(fields []string, values []interface{}, suspended bool, students *[]models.Student) (err error)
	RegisterStudent(tch *models.Teacher, stds *[]models.Student) (err error)
	GetCommonStudents(fields []string, values []interface{}, students *[]models.Student) (err error)
	SuspendStudent(email string) (err error)
	RetrieveNotifications(teacherEmail string, student *[]models.Student) (err error)
}

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) *TeacherRepo {
	return &TeacherRepo{db: db}
}

func (t *TeacherRepo) FindTeacher(email string, tch *models.Teacher) (err error) {
	if err := t.db.Model(&models.Teacher{}).Where("email = ?", email).Find(tch).Error; err != nil {
		return errors.New("Couldn't find teacher")
	}

	return nil
}

func (t *TeacherRepo) FindStudent(fields []string, values []interface{}, students *[]models.Student) (err error) {
	if err := t.db.Model(&models.Student{}).Where(strings.Join(fields, " OR "), values...).Find(students).Error; err != nil {
		return errors.New("Couldn't find student(s)")
	}

	return nil
}

func (t *TeacherRepo) FindOneStudent(email string, student *models.Student) (err error) {
	if err := t.db.Model(&models.Student{}).Where("email = ?", email).Find(student).Error; err != nil {
		return errors.New("Couldn't find student")
	}

	return nil
}

func (t *TeacherRepo) FindStudentsForNotifications(fields []string, values []interface{}, suspended bool, students *[]models.Student) (err error) {
	if err := t.db.Model(&models.Student{}).Where(strings.Join(fields, " OR "), values...).Where("suspended = ?", suspended).Find(students).Error; err != nil {
		return errors.New("Couldn't find student(s)")
	}

	return nil
}

func (t *TeacherRepo) RegisterStudent(tch *models.Teacher, students *[]models.Student) (err error) {
	if err := t.db.Debug().Model(tch).Association("Students").Append(students); err != nil {
		return errors.New("Couldn't register student")
	}

	return nil
}

func (t *TeacherRepo) GetCommonStudents(fields []string, values []interface{}, students *[]models.Student) (err error) {
	if err := t.db.Debug().Select("DISTINCT students.email").
		Joins("JOIN registers ON registers.student_id = students.id AND registers.student_email = students.email").
		Joins("JOIN teachers ON registers.teacher_id = teachers.id AND registers.student_email = students.email").
		Where(strings.Join(fields, " OR "), values...).
		Find(students).Error; err != nil {
		return errors.New("Error when suspeding student")
	}

	return nil
}

func (t *TeacherRepo) SuspendStudent(email string) (err error) {
	if err := t.db.Model(&models.Student{}).
		Where("email = ?", email).
		Update("suspended", true).Error; err != nil {
		return errors.New("Error when suspending student")
	}

	return nil
}

func (t *TeacherRepo) RetrieveNotifications(teacherEmail string, students *[]models.Student) (err error) {
	if err := t.db.Model(&models.Student{}).
		Select("DISTINCT students.email").
		Joins("JOIN registers ON registers.student_id = students.id AND registers.student_email = students.email").
		Joins("JOIN teachers ON registers.teacher_id = teachers.id AND registers.student_email = students.email").
		Where("registers.teacher_email = ?", teacherEmail).
		Where("students.suspended = ?", false).
		Find(students).Error; err != nil {
		return errors.New("Error when retrieving notifications")
	}

	return nil
}
