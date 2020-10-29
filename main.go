package main

import (
	"fmt"
	"github.com/njilrem/go-rest-atm/config"
	"github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/routes"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func connectToDatabase() {
	for i := 0; i < 30; i++ {
		config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
		if err != nil {
			fmt.Println("Status:", err)
		} else {
			fmt.Println("Connected Successfully")
			break
		}
		time.Sleep(time.Second * 10)
	}
}

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
		connectToDatabase()
	} else {
		fmt.Printf("Connected to the DB")
	}

	config.DB.AutoMigrate(&models.Account{})

	r := routes.SetupRouter()
	//running
	r.Run()
}