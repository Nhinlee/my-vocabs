package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/words/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Word added successfully"})
	})

	r.GET("/api/v1/words/review", func(c *gin.Context) {
		words := []string{"apple", "banana", "cherry"}
		c.JSON(http.StatusOK, gin.H{"words": words})
	})

	r.POST("/api/v1/words/complete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Word marked as reviewed"})
	})

	r.Run() // Run the server on default port 8080
}
