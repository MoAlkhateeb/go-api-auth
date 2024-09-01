package main

import (
	"log"

	"github.com/MoAlkhateeb/go-api-auth/cmd/api"
	"github.com/MoAlkhateeb/go-api-auth/config"
	"github.com/MoAlkhateeb/go-api-auth/db"
	"gorm.io/gorm"
)

func main() {
	db := db.NewSQLiteStorage(config.Envs.DBName, gorm.Config{})
	server := api.NewAPIServer(config.Envs.PublicHost+":"+config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
