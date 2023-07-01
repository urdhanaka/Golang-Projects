package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos []ToDo
var lastID int

func main() {
	r := gin.Default()

	r.GET("/", helloWorld)
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodos)
	r.DELETE("/todos/:id", deleteTodos)

	r.Run(":8080")
}

func helloWorld(c *gin.Context) {
	c.JSON(200, "Hello World!")
}

func getTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func createTodos(c *gin.Context) {
	var todo ToDo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	lastID++
	todo.ID = lastID
	todos = append(todos, todo)

	c.JSON(201, todo)
}

func deleteTodos(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "todo not found"})
}
