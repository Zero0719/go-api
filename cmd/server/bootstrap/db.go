package bootstrap

import "go-api/internal/db"

func initDB() {
	db.Init()
}