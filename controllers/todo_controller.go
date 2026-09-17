package controllers

import (
	"strconv"
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
func GetTodos(c *gin.Context) {
	mu.Lock()
	result := make([]models.Todo, 0, len(todos))
	for _, t := range todos {
		result = append(result, t)
	}
	mu.Unlock()

	c.JSON(200, result)
}
func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	mu.Lock()
	todo, ok := todos[id]
	mu.Unlock()

	if !ok {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(200, todo)
}
