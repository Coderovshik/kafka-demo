package main

import (
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/app"
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/config"
)

func main() {
	cfg := config.New()
	app := app.New(cfg)

	app.Run()
}
