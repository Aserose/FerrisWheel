package vk

import (
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/storage/boltDB"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestVk(t *testing.T) {

	log := logger.NewLogger()

	db := boltDB.NewDB(log)
	db.InitDB()

	cfg,_,_ := config.Init()

	vkApi := NewApiVk(log)
	vkApi.InitVk(cfg.AccessKeyVK)


	ParamsTest := map[string]interface{}{
		"lat":        51.501346247072156,
		"long":       -0.14187956027359508,
		"start_time": 343547804,
		"end_time":   1448208985,
		"count":      10,
		"radius":     5000,
	}

	Convey("setup", t, func() {
		vkApi.SetParams(ParamsTest)
		Convey("get", func() {
			result, err := vkApi.Request()
			So(err, ShouldBeEmpty)
			So(result, ShouldNotBeEmpty)
		})
	})
}
