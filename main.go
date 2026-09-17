package main

import (
	"log"
	"todo-api/routes"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
