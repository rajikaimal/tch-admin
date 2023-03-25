package main

import (
	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
)

func main() {
	db.ConnectToDB()

	teacherOne := models.Teacher{Id: 1, Email: "teacher1@gmail.com", Name: "Teacher 1"}
	teacherTwo := models.Teacher{Id: 2, Email: "teacher2@gmail.com", Name: "Teacher 2"}
	teacherThree := models.Teacher{Id: 3, Email: "teacher3@gmail.com", Name: "Teacher 3"}

	db.DB.Create(&teacherOne)
	db.DB.Create(&teacherTwo)
	db.DB.Create(&teacherThree)

	studentOne := models.Student{Id: 1, Email: "student1@gmail.com", Name: "Student 1", Suspended: false}
	studentTwo := models.Student{Id: 2, Email: "student2@gmail.com", Name: "Student 2", Suspended: false}
	studentThree := models.Student{Id: 3, Email: "student3@gmail.com", Name: "Student 3", Suspended: false}

	db.DB.Create(&studentOne)
	db.DB.Create(&studentTwo)
	db.DB.Create(&studentThree)

	db.DB.Debug().Model(&teacherOne).Association("Students").Append([]models.Student{studentOne})
	db.DB.Debug().Model(&teacherTwo).Association("Students").Append([]models.Student{studentOne, studentTwo})
}

func cleanup() {
	db.DB.Delete(&models.Teacher{})
	db.DB.Delete(&models.Student{})
}
