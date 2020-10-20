package main

import (
	"github.com/njilrem/go-rest-atm/routes"
	"github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Account{})

	r := routes.SetupRouter()
	//running
	r.Run()
}