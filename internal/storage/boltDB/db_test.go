package boltDB

import (
	"github.com/Aserose/ferrisWheel/internal/logger"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	chatId = "2142121"
	example1 = "google.com"
	example2 = "fbi.com"
	testToken = "421421512r21tfw"
)

func TestDB(t *testing.T) {

	log := logger.NewLogger()

	Convey("setup", t, func() {
		db := NewDB(log)
		Convey("initialization", func() {
			err := db.InitDB()
			So(err, ShouldBeEmpty)
			Convey("put", func() {
				err := db.PutBlacklist(example1,chatId)
				So(err, ShouldBeEmpty)
				Convey("get", func() {
					result, _ := db.GetBlacklist(chatId)
					log.Info("get: getBlacklist result: ",result)
					So(result, ShouldResemble, []interface{}{example1})

					Convey("delete", func() {
						err := db.DeleteFromBlacklist(example1,chatId)
						So(err, ShouldBeEmpty)
						result, _ := db.GetBlacklist(chatId)
						So(result, ShouldBeEmpty)

						Convey("multiply crud", func() {
							db.PutBlacklist(example2,chatId)
							db.PutBlacklist(example1,chatId)

							result, _ := db.GetBlacklist(chatId)
							log.Info("multiply crud: getBlacklist result: ",result)
							So(result, ShouldResemble, []interface{}{example2, example1})

							db.DeleteFromBlacklist(example1, chatId)
							db.DeleteFromBlacklist(example2, chatId)

							resultDelete, _ := db.GetBlacklist(chatId)
							So(resultDelete, ShouldBeEmpty)

							Convey("put token", func() {
								log.Print("put token")
								err := db.PutToken(testToken,chatId)
								So(err, ShouldBeEmpty)

								Convey("get token", func() {
									access := db.GetToken(chatId)
									log.Print(access)
									So(access, ShouldResemble, map[string]interface{}{"vk": testToken})
								})
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
}
