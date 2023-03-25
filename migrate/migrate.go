package main

import (
	"fmt"

	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
)

func main() {
	db.ConnectToDB()

	db.DB.AutoMigrate(&models.Teacher{})
	db.DB.AutoMigrate(&models.Student{})

	fmt.Println("Migrations completed!")
}
