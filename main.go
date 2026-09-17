package main

import (
	"log"
	"os"
	"todo-api/routes"
)

func main() {
	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local development
	}

	log.Println("Server running on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
