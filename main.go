package main

import (
	"hilfling-oauth/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()            // Set router
	r.LoadHTMLGlob("templates/*") // Load templates

	// Setup route group for API
	api := r.Group("/")
	api.GET("/", handlers.GetRoot) // Get root structure with links to all paths

	if err := r.Run(":7070"); err != nil {
		panic(err)
	}

}
