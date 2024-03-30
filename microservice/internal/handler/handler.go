package handler

import "golang.org/x/sync/errgroup"

type ProdConService interface {
	Consume() error
	Produce() error
}

type Handler struct {
	service ProdConService
}

func New(s ProdConService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) ConsumeAndProduce() error {
	g := &errgroup.Group{}

	g.Go(h.service.Consume)
	g.Go(h.service.Produce)

	return g.Wait()
}
