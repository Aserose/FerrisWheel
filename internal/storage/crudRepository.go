package storage

func (d *database) Put(id interface{},chatId string) {
	if err := d.Db.PutBlacklist(id,chatId); err != nil {
		d.logger.Errorf("database: error entering data from the database: %s", err.Error())
	}
}

func (d *database) Get(chatId string) []interface{} {
	blacklist, err := d.Db.GetBlacklist(chatId)
	if err != nil {
		d.logger.Errorf("database: error getting data: %s", err.Error())
	}

	return blacklist
}

func (d *database) Delete(id interface{},chatId string) {
	if err := d.Db.DeleteFromBlacklist(id,chatId); err != nil {
		d.logger.Errorf("database: error deleting data: %s", err.Error())
	}
}

func (d *database) PutToken(accessToken string,chatId string) {
	if err := d.Db.PutToken(accessToken,chatId); err != nil {
		d.logger.Errorf("database: error saving token")
	}
}

func (d *database) GetToken(chatId string) map[string]interface{} {
	return d.Db.GetToken(chatId)
}

func (d *database) CreateToken(accessToken string,chatId string) map[string]interface{}{
	if err := d.Db.PutToken(accessToken,chatId); err != nil {
		d.logger.Errorf("database: error saving token")
	}
	return d.Db.GetToken(chatId)
}