package main

import (
	"demonstrasi/config"
	"demonstrasi/route"
	"demonstrasi/util"

	"github.com/go-playground/validator"
)

func main() {
	// Inisialisasi koneksi database
	config.InitDB()

	// Panggil InitialMigration untuk melakukan migrasi otomatis
	config.InitialMigration()

	e := route.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}

	// Mulai server
	e.Logger.Fatal(e.Start(":4040"))
}
