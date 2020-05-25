package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kentoje/gin-test-api/handler"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/tasks/", handler.HandleGetTasks)
	err := router.Run()
	if err != nil {
		log.Fatalf("Server error: %v", router)
	}
}

