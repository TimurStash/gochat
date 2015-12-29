package models

import(
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id     uint `gorm:"primary_key"`
	Username string `type:varchar(45);unique`
	Password string `sql:"size:45"`
}

func (u User) SecureData() User  {
	u.Password = ""
	return u
}

func (u *User) HashPassword(){
	hasher := md5.New()
	hasher.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hasher.Sum(nil))

}