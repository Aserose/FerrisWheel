package web

import (
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWebClient(t *testing.T) {

	log := logger.NewLogger()
	cfg, cfgServer, _ := config.Init()

	Convey("setup", t, func() {
		client := NewClient(log, cfg, cfgServer)
		err := client.Initialization()
		Convey("initialization", func() {
			So(err, ShouldBeEmpty)
		})
		Reset(func() {
			t.Log("finish")
		})
	})
}
