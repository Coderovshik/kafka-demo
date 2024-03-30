package router

import (
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func New(h *handler.Handler) *Router {
	router := gin.Default()
	router.GET("/ping", h.Ping)

	return &Router{
		engine: router,
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
