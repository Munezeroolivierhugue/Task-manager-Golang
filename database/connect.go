package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Munezeroolivierhugue/Task-manager-Golang/models"
)

var DB *gorm.DB

func Connect(){
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to database", err)
	}

	DB = db
	fmt.Println("Database connected successfully")

	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("failed to migrate the database", err)
	}

	fmt.Println("Migration completed")
}