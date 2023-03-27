package main

import (
	"github.com/rajikaimal/tch-admin/config"
	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
)

func main() {
	config := config.InitConfig()
	db := db.ConnectToDB(config.DB)

	teacherOne := models.Teacher{Id: 1, Email: "teacher1@gmail.com", Name: "Teacher 1"}
	teacherTwo := models.Teacher{Id: 2, Email: "teacher2@gmail.com", Name: "Teacher 2"}
	teacherThree := models.Teacher{Id: 3, Email: "teacher3@gmail.com", Name: "Teacher 3"}

	db.Create(&teacherOne)
	db.Create(&teacherTwo)
	db.Create(&teacherThree)

	studentOne := models.Student{Id: 1, Email: "student1@gmail.com", Name: "Student 1", Suspended: false}
	studentTwo := models.Student{Id: 2, Email: "student2@gmail.com", Name: "Student 2", Suspended: false}
	studentThree := models.Student{Id: 3, Email: "student3@gmail.com", Name: "Student 3", Suspended: false}

	db.Create(&studentOne)
	db.Create(&studentTwo)
	db.Create(&studentThree)

	db.Debug().Model(&teacherOne).Association("Students").Append([]models.Student{studentOne})
	db.Debug().Model(&teacherTwo).Association("Students").Append([]models.Student{studentOne, studentTwo})
}

func cleanup() {
	config := config.InitConfig()
	db := db.ConnectToDB(config.DB)

	db.Delete(&models.Teacher{})
	db.Delete(&models.Student{})
}
