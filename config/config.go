package config

import (
	"demonstrasi/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Config struct {
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_PORT     string
	MYSQL_HOST     string
	MYSQL_DATABASE string
}

func InitDB() {
	config := Config{
		MYSQL_USER:     os.Getenv("MYSQL_USER"),
		MYSQL_PASSWORD: os.Getenv("MYSQL_PASSWORD"),
		MYSQL_PORT:     os.Getenv("MYSQL_PORT"),
		MYSQL_HOST:     os.Getenv("MYSQL_HOST"),
		MYSQL_DATABASE: os.Getenv("MYSQL_DATABASE"),
	}

	// Logging variabel lingkungan untuk debugging
	fmt.Printf("MYSQL_USER: %s\n", config.MYSQL_USER)
	fmt.Printf("MYSQL_PASSWORD: %s\n", config.MYSQL_PASSWORD)
	fmt.Printf("MYSQL_HOST: %s\n", config.MYSQL_HOST)
	fmt.Printf("MYSQL_PORT: %s\n", config.MYSQL_PORT)
	fmt.Printf("MYSQL_DATABASE: %s\n", config.MYSQL_DATABASE)

	// Membuat connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_DATABASE,
	)

	// Logging connection string untuk debugging
	fmt.Printf("Connection String: %s\n", connectionString)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		panic(err)
	}
}

func InitialMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.TbLogin{},
		&models.TbInformation{},
		&models.TbComment{},
		&models.TbGalery{},
		&models.TbMonument{},
	)
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
