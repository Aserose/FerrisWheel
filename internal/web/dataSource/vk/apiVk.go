package vk

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/SevereCloud/vksdk/v2/api"
)

type ApiVk interface {
	InitVk(token string) string
	Request() (map[string][]interface{}, error)
	IncomingPointsman(str string) map[string][]interface{}
	OutgoingPointsman() map[string][]interface{}
	SetParams(map[string]interface{})
}

type apiVk struct {
	Vk     *api.VK
	Params api.Params
	Photo  api.PhotosSearchResponse
	logger logger.Logger
}

func NewApiVk(log logger.Logger) ApiVk {
	return &apiVk{logger: log}
}
