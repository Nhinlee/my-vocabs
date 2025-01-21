package apis

import (
	db "my_vocabs/internal/my_vocabs/db/sqlc"
	genid "my_vocabs/pkg/gen_id"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
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

type ReviewWordsResponse struct {
	Count int      `json:"count" binding:"required"`
	Words []string `json:"words" binding:"required"`
}

func (s *Server) reviewWords(ctx *gin.Context) {
	vocabs, err := s.dbStore.ReviewVocabs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	words := make([]string, 0)
	for _, v := range vocabs {
		words = append(words, v.Word)
	}

	ctx.JSON(http.StatusOK, ReviewWordsResponse{
		Count: len(words),
		Words: words,
	})
}

type CompleteWordRequest struct {
	Word string `json:"word" binding:"required"`
}

func (s *Server) completeWord(ctx *gin.Context) {
	var req CompleteWordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	vocab, err := s.dbStore.GetVocabByName(ctx, req.Word)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Update next review time and reviewed time
	updatedReviewedTime := vocab.ReviewedTime.Int32 + 1
	_, err = s.dbStore.UpdateNextReviewByName(ctx, db.UpdateNextReviewByNameParams{
		Word: req.Word,
		NextReview: pgtype.Timestamptz{
			Time:  vocab.NextReview.Time.AddDate(0, 0, int(updatedReviewedTime)+2),
			Valid: true,
		},
		ReviewedTime: pgtype.Int4{
			Int32: updatedReviewedTime,
			Valid: true,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Word completed successfully"})
}

type DeleteWordRequest struct {
	Word string `json:"word" binding:"required"`
}

func (s *Server) deleteWord(ctx *gin.Context) {
	var req DeleteWordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := s.dbStore.DeleteVocabByName(ctx, req.Word)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Word deleted successfully"})
}

func (s *Server) listVocabs(ctx *gin.Context) {
	wordFilter := ctx.Query("word")

	vocabs, err := s.dbStore.ListVocabsByFilter(ctx, wordFilter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, vocabs)
}
