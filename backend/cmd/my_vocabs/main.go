package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// API to add a new word to the dictionary
	r.POST("/api/v1/words/add", func(c *gin.Context) {
		// Placeholder implementation: Extract word details from request
		// Save it to the database (to be implemented)
		c.JSON(http.StatusOK, gin.H{"message": "Word added successfully"})
	})

	// API to list words that need to be reviewed each day
	r.GET("/api/v1/words/review", func(c *gin.Context) {
		// Placeholder implementation: Retrieve words due for review from database
		// (to be implemented)
		c.JSON(http.StatusOK, gin.H{"message": "List of words to review"})
	})

	// API to mark a word as reviewed
	r.POST("/api/v1/words/complete", func(c *gin.Context) {
		// Placeholder implementation: Extract word ID from request
		// Mark word as reviewed in the database (to be implemented)
		c.JSON(http.StatusOK, gin.H{"message": "Word marked as reviewed"})
	})

	r.Run() // Run the server on default port 8080
}
