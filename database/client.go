package database

import (
	"log"
	"restAPI/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	DB.AutoMigrate(&model.User{})
	log.Println("Database Migration Completed!")

	if err := DB.Exec("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), email VARCHAR(255), password VARCHAR(255))").Error; err != nil {
		log.Fatal(err)
		panic("Cannot create users table")
	}
	log.Println("Database Migration Completed!")
}
