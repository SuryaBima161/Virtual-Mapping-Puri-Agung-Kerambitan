package main

import (
	"demonstrasi/config"
	"demonstrasi/route"
	"demonstrasi/util"
	"log"

	"github.com/go-playground/validator"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke database
	dsn := "root:wihpGGgSlrCkUieNGflJKzjrBWrcnPaw@tcp(monorail.proxy.rlwy.net:23451)/railway?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Successfully connected to database!")
	}

	// Panggil InitialMigration untuk melakukan migrasi otomatis
	config.InitialMigration(DB)
	log.Println("Database migration completed!")

	e := route.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}

	// Mulai server
	e.Logger.Fatal(e.Start(":4040"))
}
