package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const (
	host     = "postgres"
	user     = "user"
	password = "pass-word"
	db       = "pass-db"
	port     = "5432"
)

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, db, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect postgres: " + err.Error())
	}

	log.Println("Successfully connected to postgres")
}
