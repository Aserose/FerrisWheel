package main

import (
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/web"
)

func main() {
	log := logger.NewLogger()
	cfg, cfgServer, err := config.Init()
	if err != nil {
		err.Error()
	}
	if err := webInit(log, cfg, cfgServer); err != nil {
		log.Errorf("web initialization error: %s", err.Error())
	}
}

func webInit(log logger.Logger, cfg *config.Config, cfgServer *config.ServerConfig) error {
	webClient := web.NewClient(log, cfg, cfgServer)

	if err := webClient.Initialization(); err != nil {
		log.Errorf("client authorization error:  %s", err.Error())
	}

	return nil
}
