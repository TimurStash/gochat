package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/TimurStash/gochat/settings"

)

var DB gorm.DB = gorm.DB{}
var Error error = nil
func init(){
	settings.Init()
	stgs := settings.Get()

	DB, Error = gorm.Open("mysql", stgs.MysqlUsername + ":" + stgs.MysqlPassword + "@/" + stgs.MysqlDbname + "?charset=utf8&parseTime=True&loc=Local")
	// Disable table name's pluralization
	DB.SingularTable(true)
}