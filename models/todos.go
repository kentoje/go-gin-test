package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kentoje/gin-test-api/config"
)

func GetAllTodos(todo *[]Todo) error {
	if err := config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetOneTodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) error {
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func UpdateOneTodo(todo *Todo, id string) error {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

func DeleteOneTodo(todo *Todo, id string) error {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}