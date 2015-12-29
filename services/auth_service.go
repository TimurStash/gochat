package services

import (
//	"github.com/TimurStash/gochat/api/parameters"
	"github.com/TimurStash/gochat/core/authentication"
//	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"fmt"
	"errors"
)

func GenerateToken (userId uint) (error, string){
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(userId)
	if err != nil {
		return errors.New("Token was not generated"), ""
	} else {
		return nil, token
	}
}

func RefreshToken(req *http.Request) (error, string) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest := getToken(req)
	userId := tokenRequest.Claims["sub"]
	authBackend.BlacklistToken(tokenRequest.Raw)


	return GenerateToken(uint(userId.(float64)))
}

func BlacklistToken(req *http.Request)  {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest := getToken(req)
	fmt.Printf("%+v\n", tokenRequest)
	authBackend.BlacklistToken(tokenRequest.Raw)
}

func getToken(req *http.Request) *jwt.Token{
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, _ := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	return tokenRequest
}