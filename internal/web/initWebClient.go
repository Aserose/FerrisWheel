package web

import (
	"github.com/Aserose/ferrisWheel/internal/server"
	"github.com/Aserose/ferrisWheel/internal/storage"
	"github.com/Aserose/ferrisWheel/internal/storage/boltDB"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/geocode"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/vk"
	"github.com/Aserose/ferrisWheel/internal/web/tg"
)

func (c *client) Initialization() error {
	c.logger.Info("webClient: preparing to init DB...")

	db := boltDB.NewDB(c.logger)
	if err := db.InitDB(); err != nil {
		return err
	}

	c.logger.Info("webClient: preparing to authorize source sites...")

	storages := storage.NewDatabase(c.logger, db)
	geo := geocode.NewGeocodeApi(c.logger)
	geo.InitGeocode(c.cfg.AccessKeyOCD)
	vkApi := vk.NewApiVk(c.logger)
	sources := dataSource.NewDataSource(c.logger, vkApi, geo)

	c.logger.Info("webClient: preparing to authorize clients...")

	TgApi := tg.NewTgApi(c.logger, sources, storages, c.cfg.AuthURL)

	c.logger.Info("webclient: start server")
	servers := server.NewAuthServer(c.logger, c.cfgServer, TgApi)
	go servers.Start()

	c.logger.Info("webClient: preparing to authorize tg")

	if err := TgApi.StartTG(c.cfg.AccessKeyTG); err != nil {
		c.logger.Errorf("webClient: initialization TG error %s", err.Error())
	}

	c.logger.Info("webClient: preparing to establish a connection with tg chat")
	if err := TgApi.Connection(); err != nil {
		c.logger.Errorf("webClient: chat connection error: %s", err.Error())
	}

	return nil
}