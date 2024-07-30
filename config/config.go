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
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: os.Getenv("MYSQL_USER"),
		DB_Password: os.Getenv("MYSQL_PASSWORD"),
		DB_Port:     os.Getenv("MYSQL_PORT"),
		DB_Host:     os.Getenv("MYSQL_HOST"),
		DB_Name:     os.Getenv("MYSQL_DATABASE"),
	}

	fmt.Printf("Connecting to database with: %s:%s@tcp(%s:%s)/%s\n",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.TbLogin{}, &models.TbInformation{}, &models.TbComment{}, &models.TbGalery{}, &models.TbMonument{})
}
