package controllers

import (
	"sync"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

var (
	todos  = make(map[int]models.Todo)
	nextID = 1
	mu     sync.Mutex
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	todo.ID = nextID
	nextID++
	todos[todo.ID] = todo
	mu.Unlock()
	c.JSON(201, todo)
}
