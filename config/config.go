package config

import (
	"demonstrasi/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	// Log connection string for debugging
	fmt.Printf("Connecting to database with DSN: %s\n", dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to database!")
}

func InitialMigration() {
	if DB == nil {
		log.Fatalf("Database connection is not initialized")
	}
	err := DB.AutoMigrate(
		&models.TbLogin{},
		&models.TbInformation{},
		&models.TbComment{},
		&models.TbGalery{},
		&models.TbMonument{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}
