package geocode

import (
	"github.com/codingsince1985/geo-golang/opencage"
)

func (g *geocodeApi) InitGeocode(token string) {
	g.logger.Info("geocode: init openCage")
	g.OpenCageData = opencage.Geocoder(token)
	g.logger.Info("geocode: openCage ok")
}
