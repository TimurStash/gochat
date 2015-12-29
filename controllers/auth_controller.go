package controllers

import (
	"github.com/TimurStash/gochat/services"
	"github.com/TimurStash/gochat/services/models"
	"github.com/TimurStash/gochat/core/mysql"
	"encoding/json"
	"net/http"
	"fmt"
)

func SignUp(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&requestUser)

	//Check if the user with provided username is already exist
	mysql.DB.Where("username = ?", requestUser.Username).First(&requestUser)

	//This is new user
	if requestUser.Id == 0 {
		requestUser.HashPassword()
		mysql.DB.Create(&requestUser)
		_, token := services.GenerateToken(requestUser.Id)
		response := models.Response{Ok:true, Data: requestUser.SecureData(), Message: "User successfully created", Token: token, TokenChanged: true}
		response.SendJson(w)
		return
	}

	//THe user is already exist
	if requestUser.Id > 0 {
		response := models.Response{Ok:false, Data: nil, Message: "User already exist", Token:"", TokenChanged: false}
		response.SendJson(w)
		return
	}

}



func Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&requestUser)
	fmt.Printf("%+v\n", requestUser)
	requestUser.HashPassword()
	mysql.DB.Where("username = ? AND password = ?", requestUser.Username, requestUser.Password).First(&requestUser)
	if requestUser.Id > 0 {
		_, token := services.GenerateToken(requestUser.Id)
		response := models.Response{Ok:true, Data: requestUser.SecureData(), Message: "User successfully logged in", Token: token, TokenChanged: true}
		response.SendJson(w)
	} else {
		response := models.Response{Ok:false, Data: nil, Message: "User was not found", Token:"", TokenChanged: false}
		response.SendJson(w)
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	_, token := services.RefreshToken(r)
	response := models.Response{Ok:true, Data: nil, Message: "Token changed", Token:token, TokenChanged: true}
	response.SendJson(w)
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	services.BlacklistToken(r)
	response := models.Response{Ok:true, Data: nil, Message: "User successfully logged out", Token:"", TokenChanged: false}
	response.SendJson(w)

}