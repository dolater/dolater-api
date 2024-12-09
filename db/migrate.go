package db

import (
	"log"

	"github.com/kantacky/p2hacks2024-test-api/model"
)

func Migrate() {
	db, err := GormDB("public")
	if err != nil {
		log.Println("Failed to connect database")
	}
	defer func() {
		sqldb, err := db.DB()
		if err != nil {
			log.Println("Failed to close database connection")
		}
		sqldb.Close()
	}()

	err = db.AutoMigrate(&model.FCMToken{})
	if err != nil {
		log.Println("Failed to migrate database")
	}
}
