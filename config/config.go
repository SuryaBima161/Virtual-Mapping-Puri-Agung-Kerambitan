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

	fmt.Printf("Connecting to database with: %s:%s@tcp(%s:%s)/%s\n",
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
		config.MYSQL_PORT,
		config.MYSQL_HOST,
		config.MYSQL_DATABASE,
	)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
		config.MYSQL_PORT,
		config.MYSQL_HOST,
		config.MYSQL_DATABASE,
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
