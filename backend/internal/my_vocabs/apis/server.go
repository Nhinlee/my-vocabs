package apis

import (
	"my_vocabs/internal/my_vocabs/config"
	db "my_vocabs/internal/my_vocabs/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg     *config.Config
	router  *gin.Engine
	dbStore db.Store
}

func NewServer(config *config.Config, dbStore db.Store) (*Server, error) {
	s := &Server{
		cfg:     config,
		dbStore: dbStore,
	}

	s.SetupRouter()

	return s, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	router.POST("/api/v1/vocabs/add", s.newVocab)
	router.GET("/api/v1/vocabs/review", s.reviewWords)
	router.POST("/api/v1/vocabs/complete", s.completeWord)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
