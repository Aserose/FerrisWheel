package geocode

import (
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	defaultAddress = "red square"
	wrongAddress   = "xzvxcbwa"
)

func TestGeocode(t *testing.T) {

	cfg,_,_ := config.Init()

	log := logger.NewLogger()

	g := NewGeocodeApi(log)

	Convey("setup", t, func() {
		g.InitGeocode(cfg.AccessKeyOCD)
		_, err := g.GetCoordinates(defaultAddress)
		Convey("getCoordinates", func() {
			So(err, ShouldBeEmpty)

			Convey("not found", func() {
				coordinates, _ := g.GetCoordinates(wrongAddress)

				So(coordinates, ShouldBeNil)
			})
		})
		Reset(func() {
			t.Log("finish")
		})
	})
}
