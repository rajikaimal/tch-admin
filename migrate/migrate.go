package main

import (
	"fmt"

	"github.com/rajikaimal/tch-admin/config"
	"github.com/rajikaimal/tch-admin/db"
	"github.com/rajikaimal/tch-admin/models"
)

func main() {
	config := config.InitConfig()
	db := db.ConnectToDB(config.DB)

	db.AutoMigrate(&models.Teacher{})
	db.AutoMigrate(&models.Student{})

	fmt.Println("Migrations completed!")
}
