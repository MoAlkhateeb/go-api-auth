package main

import (
	"log"

	"github.com/MoAlkhateeb/go-api-auth/config"
	"github.com/MoAlkhateeb/go-api-auth/db"
	"github.com/MoAlkhateeb/go-api-auth/types"
	"gorm.io/gorm"
)

func main() {
	db := db.NewSQLiteStorage(config.Envs.DBName, gorm.Config{})
	db.AutoMigrate(&types.User{})
	log.Println("Automigrated Successfully.")
}
