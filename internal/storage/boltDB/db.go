package boltDB

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/boltdb/bolt"
)

type DB interface {
	PutBlacklist(Id interface{},chatId string) error
	GetBlacklist(chatId string) ([]interface{}, error)
	DeleteFromBlacklist(id interface{},chatId string) error
	PutToken(accessToken interface{},chatId string) error
	GetToken(chatId string) map[string]interface{}
	InitDB() error
}

type db struct {
	Db     *bolt.DB
	logger logger.Logger
}

func NewDB(log logger.Logger) DB {
	return &db{
		logger: log,
	}
}
