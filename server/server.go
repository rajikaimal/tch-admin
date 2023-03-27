package server

import (
	"github.com/rajikaimal/tch-admin/config"
	"github.com/rajikaimal/tch-admin/db"
)

func Start() {
	// read config
	config := config.InitConfig()
	// connect to db
	db := db.ConnectToDB(config.DB)
	// initialize router
	r := NewRouter(db)

	r.Run(":8080")
}
