package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kentoje/gin-test-api/models"
	"net/http"
)

func HandleGetTasks(c *gin.Context) {
	var tasks []models.Task

	oneTask := models.Task{
		Title: "Hello",
		Body:  "I like eating a lot!",
	}

	tasks = append(tasks, oneTask)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
