package controller

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	s := &Server{}
	s.setupRouter()
	return s
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.POST("/validate", s.validateBeneficiary)
	s.router = router
}

func (s *Server) Start(address string) {
	s.router.Run(address)
}
