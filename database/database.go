package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Instance - database instance
var Instance *gorm.DB

// Connect - handles connection to the database
func Connect() (db *gorm.DB, err error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := 5432
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword,
	)
	
	db, err = gorm.Open("postgres", dbInfo)
	Instance = db
	return db, err
}