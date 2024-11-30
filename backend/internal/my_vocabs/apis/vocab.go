package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) addWord(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Word added successfully"})
}

func (s *Server) reviewWords(ctx *gin.Context) {
	words := []string{"apple", "banana", "cherry"}
	ctx.JSON(http.StatusOK, gin.H{"words": words})
}

func (s *Server) completeWord(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Word marked as reviewed"})
}
