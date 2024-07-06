package main

import (
	"demonstrasi/config"
	"demonstrasi/route"
	"demonstrasi/util"

	"github.com/go-playground/validator"
)

// import "golang.org/x/net/route"
func init() {
	config.InitDB()
	config.InitialMigration()
	// seeder.DBSeed(config.DB)
}

func main() {
	e := route.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":8000"))
}
