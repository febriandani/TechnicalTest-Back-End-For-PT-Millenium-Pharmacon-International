package database

import (
	"fmt"
	"log"
	"technical-test-02-10-22/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "junior34"
	dbPort   = "5432"
	dbName   = "technical_test"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("sukses connect to db")
	db.Debug().AutoMigrate(models.File{})
}

func GetDB() *gorm.DB {
	return db
}
