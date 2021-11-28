package tg

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/storage"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgApi interface {
	StartTG(token string) error
	Connection() error
	CreateAccessKey(accessKey string)
}

type (
	tgApi struct {
		bot          *tgbotapi.BotAPI
		updateConfig tgbotapi.UpdateConfig
		updatesChan  tgbotapi.UpdatesChannel
		result       map[string][]interface{}
		blacklist    []interface{}
		DataCollect  dataSource.DataSources
		Storage      storage.Database
		logger       logger.Logger
		authURL      string
		chatId string
		sC serviceConstruction
	}

	serviceConstruction struct {
		indexCounter int
		removeIsOn   bool
		requestIsOn  bool
	}
)

func NewTgApi(log logger.Logger, sources dataSource.DataSources, database storage.Database, authURL string) TgApi {
	return &tgApi{logger: log,
		DataCollect: sources,
		Storage:     database,
		authURL:     authURL}
}
