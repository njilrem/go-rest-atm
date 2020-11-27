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

	err = config.DB.AutoMigrate(&models.Account{})

	if err != nil {
		fmt.Printf("Accounts auto migration failed")
	}

	err = config.DB.AutoMigrate(&models.Card{})

	if err != nil {
		fmt.Printf("Cards auto migration failed")
	}

	err = config.DB.AutoMigrate(&models.Transaction{})

	if err != nil {
		fmt.Printf("Transactions auto migration failed")
	}

	r := routes.SetupRouter()
	//running
	err = r.Run()

	if err != nil {
		panic("SERVER DIDN'T START")
	}
}