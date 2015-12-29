package models

type Token struct {
	Id     uint `gorm:"primary_key"`
	Token string `type:TEXT`
}
