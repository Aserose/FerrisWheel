package boltDB

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func (d *db) PutBlacklist(Id interface{},chatId string) error {
	d.createBlacklist(chatId)

	err := d.Db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("blacklist")).Bucket([]byte(chatId)).Put([]byte(fmt.Sprintf("%v", Id)), []byte(fmt.Sprintf("%v", Id)))
		if err != nil {
			d.logger.Errorf("bolt: could not insert id: %s", err.Error())
		}
		return nil
	})

	return err
}

func (d *db) GetBlacklist(chatId string) ([]interface{}, error) {
	d.createBlacklist(chatId)

	var blacklist []interface{}
	if err := d.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("blacklist")).Bucket([]byte(chatId))
		b.ForEach(func(k, v []byte) error {
			blacklist = append(blacklist, string(v))
			return nil
		})
		return nil
	}); err != nil {
		d.logger.Errorf("bolt: failed with %s", err.Error())
	}
	return blacklist, nil
}

func (d *db) DeleteFromBlacklist(Id interface{},chatId string) error {
	if err := d.Db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("blacklist")).Bucket([]byte(chatId)).Delete([]byte(fmt.Sprintf("%v", Id)))
		if err != nil {
			d.logger.Errorf("bolt: error deleting data: %s", err.Error())
		}
		return nil
	}); err != nil {
		d.logger.Errorf("bolt: failed with %s", err.Error())
	}
	return nil
}

func (d *db) PutToken(accessToken interface{}, chatId string) error {
	if err := d.Db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("token")).Put([]byte(chatId), []byte(fmt.Sprintf("%v", accessToken)))
		if err != nil {
			d.logger.Errorf("bolt: error putting token: %s", err.Error())
		}
		return nil
	}); err != nil {
		d.logger.Errorf("bolt: failed with %s", err.Error())
	}
	return nil
}

func (d *db) GetToken(chatId string) map[string]interface{} {
	res := make(map[string]interface{})

	if err := d.Db.View(func(tx *bolt.Tx) error {
		res["vk"] = string(tx.Bucket([]byte("DB")).Bucket([]byte("token")).Get([]byte(chatId)))
		return nil
	}); err != nil {
		d.logger.Errorf("bolt: failed with %s", err.Error())
	}

	return res
}
