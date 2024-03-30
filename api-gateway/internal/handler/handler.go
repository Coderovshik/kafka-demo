package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Coderovshik/kafka-demo/api-gateway/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProdConService interface {
	Produce(ctx context.Context, message *domain.Message) error
	Consume(ctx context.Context) (*domain.Message, error)
}

type Handler struct {
	service ProdConService
}

func New(s ProdConService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	req := &domain.Message{
		ID:    uuid.New().String(),
		Name:  "Ping",
		Value: "Pong",
	}

	err := h.service.Produce(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to produce message"})
		return
	}

	res, err := h.service.Consume(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to consume message"})
		return
	}

	c.JSON(http.StatusOK, res)
}
