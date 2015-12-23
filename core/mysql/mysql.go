package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/TimurStash/gochat/settings"
)

var DB gorm.DB = nil
var Error error = nil
func init() (gorm.DB, error){

	stgs := settings.Get()

	DB, Error = gorm.Open("mysql", stgs.MysqlUsername + ":" + stgs.MysqlPassword + "@/" + stgs.MysqlDbname + "?charset=utf8&parseTime=True&loc=Local")
	return DB, Error
}