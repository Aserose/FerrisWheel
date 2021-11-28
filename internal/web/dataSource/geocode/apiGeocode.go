package geocode

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/codingsince1985/geo-golang"
)

type Geocode interface {
	InitGeocode(token string)
	GetCoordinates(address string) (map[string]interface{}, string)
}

type geocodeApi struct {
	OpenCageData geo.Geocoder
	logger       logger.Logger
}

func NewGeocodeApi(log logger.Logger) Geocode {
	return &geocodeApi{
		logger: log,
	}
}
