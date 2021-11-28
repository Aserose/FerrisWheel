package storage

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/storage/boltDB"
)

type Database interface {
	Put(id interface{},chatId string)
	Get(chatId string) []interface{}
	Delete(id interface{},chatId string)
	PutToken(accessToken string, chatId string)
	GetToken(chatId string) map[string]interface{}
	CreateToken(accessToken string,chatId string) map[string]interface{}
}

type database struct {
	Db     boltDB.DB
	TgChatId string
	logger logger.Logger
}

func NewDatabase(log logger.Logger, db boltDB.DB) Database {
	return &database{
		logger: log,
		Db:     db}
}
