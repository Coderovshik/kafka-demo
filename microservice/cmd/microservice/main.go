package main

import (
	"github.com/Coderovshik/kafka-demo/microservice/internal/app"
	"github.com/Coderovshik/kafka-demo/microservice/internal/config"
)

func main() {
	cfg := config.New()
	app := app.New(cfg)

	app.Run()
}
