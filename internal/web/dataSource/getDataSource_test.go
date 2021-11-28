package dataSource

import (
	"fmt"
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/storage/boltDB"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/geocode"
	"github.com/Aserose/ferrisWheel/internal/web/dataSource/vk"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

const (
	testChatId = "421421"
	request          = "new request"
	send             = "send request"
	defaultAddress   = "red square"
	defaultEndTime   = "10.12.2014"
	defaultStartTime = "10.12.2013"
)

func TestGetDataSource(t *testing.T) {
	log := logger.NewLogger()

	db := boltDB.NewDB(log)
	db.InitDB()


	cfg,_,_ := config.Init()


	Convey("setup", t, func() {
		vkApi := vk.NewApiVk(log)
		vkApi.InitVk(os.Getenv(cfg.AccessKeyVK))
		g := geocode.NewGeocodeApi(log)
		g.InitGeocode(cfg.AccessKeyOCD)
		dS := NewDataSource(log, vkApi, g)

		Convey("set request", func() {
			status, statusFlag := dS.SetRequestParametrs(request)
			So(status, ShouldEqual, "enter address")
			So(statusFlag, ShouldBeFalse)

			Convey("set address", func() {
				status, statusFlag = dS.SetRequestParametrs(defaultAddress)
				So(status, ShouldEqual, fmt.Sprintf("address: %s\n\n%s\nenter the end date (mm/dd/yy)\n\nor send request", defaultAddress, inputTime))
				So(statusFlag, ShouldBeFalse)

				Convey("set endTime", func() {
					status, statusFlag = dS.SetRequestParametrs(defaultEndTime)
					So(status, ShouldEqual, fmt.Sprintf("end date: %s\n\n%s\nenter the start date (mm/dd/yy) \n\nor send request", defaultEndTime, inputTime))
					So(statusFlag, ShouldBeFalse)

					Convey("set startTime", func() {
						status, statusFlag = dS.SetRequestParametrs(defaultStartTime)
						So(status, ShouldEqual, " ")
						So(statusFlag, ShouldBeTrue)

						Convey("get data request", func() {
							result := dS.GetData(request)
							So(result, ShouldNotResemble, map[string][]interface{}{})
						})
					})

					Convey("data request with endTime", func() {
						status, statusFlag = dS.SetRequestParametrs(send)
						So(status, ShouldEqual, " ")
						So(statusFlag, ShouldBeTrue)

						result := dS.GetData(request)
						So(result, ShouldNotResemble, map[string][]interface{}{})
					})
				})

				Convey("set wrong endTime", func() {
					status, statusFlag = dS.SetRequestParametrs("2515faw3 3525")
					So(status, ShouldEqual, fmt.Sprint(status))
					So(statusFlag, ShouldBeFalse)
				})

				Convey("data request with address", func() {
					status, statusFlag = dS.SetRequestParametrs(send)
					So(status, ShouldEqual, " ")
					So(statusFlag, ShouldBeTrue)

					result := dS.GetData(request)
					So(result, ShouldNotResemble, map[string][]interface{}{})

				})
			})

			Convey("set wrong address", func() {
				status, statusFlag = dS.SetRequestParametrs("gwag4242fw")
				So(status, ShouldEqual, "address not found")
				So(statusFlag, ShouldBeFalse)
			})
		})
		Reset(func() {
			t.Log("finish")
		})
	})
}
