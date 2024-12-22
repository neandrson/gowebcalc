package main

import (
	"github.com/AtariOverlord09/gowebcalc/config"
	"github.com/AtariOverlord09/gowebcalc/internal/application"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		return
	}
	defer logger.Sync()

	app := application.New(&cfg, logger)
	app.RunServer()
}
