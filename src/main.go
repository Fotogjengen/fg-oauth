package main

import (
	"fmt"
	"hilfling-oauth/database"
	"hilfling-oauth/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()            // Set router
	r.LoadHTMLGlob("templates/*") // Load templates

	// Setup route group for API
	api := r.Group("/")
	api.GET("/", handlers.GetRoot) // Get root structure with links to all paths
	api.GET("/security_levels", handlers.GetSecurityLevels)

	database.Connect()
	if err := database.DB.Ping(); err != nil {
		panic(err)
	}
	defer database.DB.Close()

	fmt.Println("Successfully connected to DB!")

	if err := r.Run(":7070"); err != nil {
		panic(err)
	}

}
