package tg

import (
	"fmt"
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/storage"
	"github.com/Aserose/ferrisWheel/internal/storage/boltDB"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/geocode"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/vk"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	defaultAddress   = "red square"
	defaultEndTime   = "10.12.2014"
	defaultStartTime = "10.12.2013"
	inputTime        = "setting the search time interval: "
	testChatId = "4123123"
)

func TestCallTG(t *testing.T) {

	cfg,_,_ := config.Init()

	Convey("setup", t, func() {
		logg := logger.NewLogger()
		db := boltDB.NewDB(logg)
		db.InitDB()

		storages := storage.NewDatabase(logg, db)

		vkApi := vk.NewApiVk(logg)
		vkApi.InitVk(cfg.AccessKeyVK)
		g := geocode.NewGeocodeApi(logg)
		g.InitGeocode(cfg.AccessKeyOCD)
		dS := dataSource.NewDataSource(logg, vkApi, g)

		tg := tgApi{
			logger:      logg,
			Storage:     storages,
			DataCollect: dS}

		Convey("new request to data sources", func() {
			status, statusFlag := tg.requestToDataSource(request)
			So(status, ShouldEqual, "enter address")
			So(statusFlag, ShouldBeFalse)

			Convey("set address", func() {
				status, statusFlag = tg.requestToDataSource(defaultAddress)
				So(status, ShouldEqual, fmt.Sprintf("address: %s\n\n%s\nenter the end date (mm/dd/yy)\n\nor send request", defaultAddress, inputTime))
				So(statusFlag, ShouldBeFalse)

				Convey("set endTime", func() {
					status, statusFlag = tg.requestToDataSource(defaultEndTime)
					So(status, ShouldEqual, fmt.Sprintf("end date: %s\n\n%s\nenter the start date (mm/dd/yy) \n\nor send request", defaultEndTime, inputTime))
					So(statusFlag, ShouldBeFalse)

					Convey("set startTime", func() {
						status, statusFlag = tg.requestToDataSource(defaultStartTime)
						So(status, ShouldEqual, "next")
						So(statusFlag, ShouldBeTrue)

						Convey("get data", func() {
							img, usr := tg.formattingMsgResponse()
							So(img, ShouldNotResemble, []interface{}{})
							So(usr, ShouldNotEqual, end)

							Convey("put blacklist", func() {
								ok := tg.inputBlacklist("1",testChatId)
								So(ok, ShouldContainSubstring, "added to blacklist")
								ok = tg.inputBlacklist("2",testChatId)
								So(ok, ShouldContainSubstring, "added to blacklist")

								Convey("remove from blacklist", func() {
									for i := 0; i <= 1; i++ {
										ok = tg.handleBlacklist(fmt.Sprint("remove"))
										So(ok, ShouldContainSubstring, "who should be removed")
										ok = tg.handleBlacklist("1")
										So(ok, ShouldContainSubstring, "has been removed")
									}
								})
							})
						})
					})
				})
			})
			Reset(func() {
				t.Log("finish")
			})
		})
	})
}
