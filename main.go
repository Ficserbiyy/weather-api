package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin router with default middleware
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server listening on 0.0.0.0:8080")
	if err := r.Run(); err != nil {
		log.Fatalf("unable to listen to server: %v", err)
	}
}
