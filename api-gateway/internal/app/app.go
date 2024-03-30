package app

import (
	"log"

	"github.com/Coderovshik/kafka-demo/api-gateway/internal/config"
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/handler"
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/router"
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/service"
)

type App struct {
	cfg    *config.Config
	router *router.Router
}

func New(cfg *config.Config) *App {
	service := service.New(cfg)
	handler := handler.New(service)
	router := router.New(handler)

	return &App{
		cfg:    cfg,
		router: router,
	}
}

func (a *App) Run() {
	log.Printf("INFO: server running %s", a.cfg.HTTPServer.Address())
	if err := a.router.Run(a.cfg.HTTPServer.Address()); err != nil {
		log.Fatal(err)
	}
}
