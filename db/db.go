package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteStorage(dbname string, cfg gorm.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbname), &cfg)
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB: Successfully Connected.")
	return db
}
