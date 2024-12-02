package apis

import (
	db "my_vocabs/internal/my_vocabs/db/sqlc"
	genid "my_vocabs/pkg/gen_id"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewVocabRequest struct {
	Word string `json:"word" binding:"required"`
}

func (s *Server) newVocab(ctx *gin.Context) {
	var req NewVocabRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	w, err := s.dbStore.CreateVocab(ctx, db.CreateVocabParams{
		VocabID: genid.NewULID(),
		Word:    req.Word,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, w)
}

func (s *Server) reviewWords(ctx *gin.Context) {
	words := []string{"apple", "banana", "cherry"}
	ctx.JSON(http.StatusOK, gin.H{"words": words})
}

func (s *Server) completeWord(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Word marked as reviewed"})
}
