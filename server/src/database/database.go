package database

import (
	"log"
	"menagment-app-2/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func New() *Postgres {

	connStr := "host=" + config.AppConfig.DBHost + " user=" + config.AppConfig.DBUser + " password=" + config.AppConfig.DBPass +
		" dbname=" + config.AppConfig.DBName + " port=" + config.AppConfig.DBPort + " sslmode=disable TimeZone=Europe/Warsaw"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &Postgres{
		db: db,
	}
}
