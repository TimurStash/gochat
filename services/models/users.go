package models

type User struct {
	id     uint `gorm:"primary_key"`
	username string `type:varchar(45);unique`
	password string `sql:"size:45"`
}