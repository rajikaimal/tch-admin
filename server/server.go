package server

import "github.com/rajikaimal/tch-admin/db"

func Start() {
	// connect to db
	db.ConnectToDB()
	// initialize router
	r := NewRouter()

	r.Run("localhost:8083")
}
