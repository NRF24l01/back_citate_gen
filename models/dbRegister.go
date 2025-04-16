package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func RegisterPostgres() (*gorm.DB){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"), os.Getenv("POSTGRES_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to get db instance: %v", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

	db.AutoMigrate(&User{})

	return db
}