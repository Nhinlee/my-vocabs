package apis

import (
	"my_vocabs/internal/my_vocabs/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	// TODO: DB store
}

func NewServer(config *config.Config) (*Server, error) {
	s := &Server{
		cfg: config,
	}

	s.SetupRouter()

	return s, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	router.POST("/api/v1/words/add", s.addWord)
	router.GET("/api/v1/words/review", s.reviewWords)
	router.POST("/api/v1/words/complete", s.completeWord)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
