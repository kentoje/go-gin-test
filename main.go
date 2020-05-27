package main

import (
	"github.com/jinzhu/gorm"
	"github.com/kentoje/gin-test-api/config"
	"github.com/kentoje/gin-test-api/models"
	"github.com/kentoje/gin-test-api/routes"
	"log"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Fatal(err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
