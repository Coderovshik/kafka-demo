package app

import (
	"log"

	"github.com/Coderovshik/kafka-demo/microservice/internal/config"
	"github.com/Coderovshik/kafka-demo/microservice/internal/handler"
	"github.com/Coderovshik/kafka-demo/microservice/internal/service"
)

type App struct {
	handler *handler.Handler
}

func New(cfg *config.Config) *App {
	service := service.New(cfg)
	handler := handler.New(service)

	return &App{
		handler: handler,
	}
}

func (a *App) Run() {
	if err := a.handler.ConsumeAndProduce(); err != nil {
		log.Fatal(err)
	}
}
