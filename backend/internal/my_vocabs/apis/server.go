package apis

import (
	"my_vocabs/internal/my_vocabs/config"
	db "my_vocabs/internal/my_vocabs/db/sqlc"
	fs "my_vocabs/pkg/file_store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine

	dbStore   db.Store
	fileStore fs.FileStore
}

func NewServer(
	config *config.Config,
	dbStore db.Store,
	fileStore fs.FileStore,
) (*Server, error) {
	s := &Server{
		cfg:       config,
		dbStore:   dbStore,
		fileStore: fileStore,
	}

	s.SetupRouter()

	return s, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/api/v1/vocabs/list", s.listVocabs)
	router.GET("/api/v1/vocabs/:id", s.getVocabByID)
	router.POST("/api/v1/vocabs/add", s.newVocab)
	router.GET("/api/v1/vocabs/review", s.reviewWords)
	router.POST("/api/v1/vocabs/complete", s.completeWord)
	router.POST("/api/v1/vocabs/delete", s.deleteWord)

	router.POST("/api/v1/upload/generate-presigned-url", s.handleGeneratePresignedURL)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
