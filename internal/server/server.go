package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start()
}

type server struct {
	handlers Handlers
}

func New(handlers Handlers) Server {
	return &server{
		handlers: handlers,
	}
}

func (s *server) Start() {
	r := gin.Default()
	applyRouter(r, s.handlers)

	r.Run()
}

func applyRouter(r *gin.Engine, h Handlers) {
	r.GET("/product", h.GetAllProducts)
}
