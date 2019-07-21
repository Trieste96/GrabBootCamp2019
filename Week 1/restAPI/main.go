package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

var db *gorm.DB

func main() {
	var err error
	log.SetOutput(os.Stdout)
	db, err := gorm.Open("mysql", "root@tcp(172.18.2.1:3306)/todolist")
	if err != nil {
		log.Fatal(fmt.Sprintf("cannot connect to database, error: %s", err))
	}

	db.LogMode(true)
	defer db.Close()

	err = db.AutoMigrate(Todo{}).Error
	if err != nil {
		log.Fatal("cannot migrate table Todo, err: ", err)
	}

	router := gin.Default()

	router.GET("/todos", listTodos)
	router.POST("/todos", createTodo)
	router.Run(":8080")

}

func listTodos(c *gin.Context) {
	var todos []Todo
	err := db.Find(todos).Error

	if err != nil {
		c.JSON(500, "cannot list todos")
	}

	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var argument struct {
		Title string
	}

	err := c.BindJSON(&argument).Error
	if err != nil {
		c.String(400, "invalid argument, err: ", err)
		return
	}

	todo := Todo{
		Title: argument.Title,
	}

	err = db.Create(&todo).Error.Error
	if err != nil {
		c.String(500, "cannot insert todo to database, err %s", err)
	}

	c.JSON(200, todo)
}
