package main

import (
	"fmt"
	"github.com/njilrem/go-rest-atm/config"
	"github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}

	config.DB.AutoMigrate(&models.Account{})

	r := routes.SetupRouter()
	//running
	r.Run()
}