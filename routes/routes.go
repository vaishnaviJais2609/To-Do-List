package routes

import (
	"net/http"
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Todo API is running"})
	})

	todos := router.Group("/todos")
	{
		todos.POST("", controllers.CreateTodo)
		todos.GET("", controllers.GetTodos)
		todos.GET("/:id", controllers.GetTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}

	return router
}
