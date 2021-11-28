package boltDB

import (
	"github.com/boltdb/bolt"
)

func (d *db) InitDB() error {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		d.logger.Errorf("bolt: could not open db, %s", err.Error())
	}

	err = db.Batch(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			d.logger.Errorf("bolt: could not create root bucket: %s", err.Error())
		}
		_, err = root.CreateBucketIfNotExists([]byte("blacklist"))
		if err != nil {
			d.logger.Errorf("bolt: could not create blacklist bucket: %s", err.Error())
		}
		_, err = root.CreateBucketIfNotExists([]byte("token"))
		if err != nil {
			d.logger.Errorf("bolt: could not create token bucket: %s", err.Error())
		}
		return nil
	})
	if err != nil {
		d.logger.Errorf("bolt: could not set up buckets, %s", err.Error())
	}
	d.logger.Info("bolt: DB Setup Done")

	d.Db = db
	return nil
}

func (d *db) createBlacklist(chatId string) {
	if err := d.Db.Update(func(tx *bolt.Tx) error {
		_,err := tx.Bucket([]byte("DB")).Bucket([]byte("blacklist")).CreateBucketIfNotExists([]byte(chatId)); if err != nil {
			d.logger.Errorf("bolt: blacklist creation error: %s", err.Error())
		}
		return nil
	}); err != nil {
		d.logger.Errorf("bolt: blacklist creation error: %s ", err.Error())
	}
}