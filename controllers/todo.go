package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kentoje/gin-test-api/models"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todo []models.Todo

	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo

	err := c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}

	err = models.CreateTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func GetATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.GetOneTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.GetOneTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}

	err = models.UpdateOneTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.DeleteOneTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"idDeleted": id})
}
