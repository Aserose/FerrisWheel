package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (tA *tgApi) StartTG(token string) error {
	tA.logger.Info("authorization TG")

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		tA.logger.Errorf("TG: authorization TG error: %s", err.Error())
	}

	tA.bot = botApi

	tA.bot.Debug = true

	tA.logger.Infof("tg: Authorized on account %s", tA.bot.Self.UserName)

	tA.updateConfig = tgbotapi.NewUpdate(0)
	tA.updateConfig.Timeout = 60

	tA.updatesChan = tA.bot.GetUpdatesChan(tA.updateConfig)
	return nil
}

func (tA *tgApi) CreateAccessKey(accessKey string) {
	tokens := tA.Storage.CreateToken(accessKey,tA.chatId)
	tA.DataCollect.InitServices(tokens)
}

func (tA *tgApi) initSources(chatId int64) {
	tA.chatId = strconv.FormatInt(chatId, 10)
	tokens := tA.Storage.GetToken(tA.chatId)
	tA.DataCollect.InitServices(tokens)
}