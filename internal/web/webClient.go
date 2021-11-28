package web

import (
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
)

type Client interface {
	Initialization() error
}

type client struct {
	logger    logger.Logger
	cfg       *config.Config
	cfgServer *config.ServerConfig
}

func NewClient(log logger.Logger, cfg *config.Config, cfgServer *config.ServerConfig) Client {
	return &client{
		logger:    log,
		cfg:       cfg,
		cfgServer: cfgServer,
	}
}
