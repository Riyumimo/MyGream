package database

import (
	"MyGream/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "ilham123"
	dbPort   = "5432"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDb() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	dsn := config

	// import c
	// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to databse", err)
	}
	fmt.Println("Sukses conecting to database")

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

}

func GetDB() *gorm.DB {
	return db
}
