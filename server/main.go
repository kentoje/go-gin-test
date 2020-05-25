package main

import (
	"github.com/gin-gonic/gin"
	"gin-test-api/models/task"
	"log"
	"net/http"
)

func handleGetTasks(c *gin.Context) {
	var tasks []task.Task

	oneTask := task.Task{
		Title: "Hello",
		Body:  "I like eating a lot!",
	}

	tasks = append(tasks, oneTask)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func main() {
	router := gin.Default()
	router.GET("/tasks/", handleGetTasks)
	err := router.Run()
	if err != nil {
		log.Fatalf("Server error: %v", router)
	}
}
