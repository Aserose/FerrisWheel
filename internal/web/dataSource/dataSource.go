package dataSource

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/geocode"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/vk"
)

type DataSources interface {
	GetData(derive string) map[string][]interface{}
	SetRequestParametrs(derive string) (string, bool)
	InitServices(tokens map[string]interface{})
	CheckStatus() string
}

type dataSource struct {
	AuthorizationStatus map[string]string
	Geocode             geocode.Geocode
	Params              map[string]interface{}
	Vk                  vk.ApiVk
	logger              logger.Logger
}

func NewDataSource(log logger.Logger, apiVk vk.ApiVk, apiGeocode geocode.Geocode) DataSources {
	return &dataSource{
		logger:  log,
		Vk:      apiVk,
		Geocode: apiGeocode,
	}
}
