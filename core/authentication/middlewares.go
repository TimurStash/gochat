package authentication

import (
	"fmt"
	jwt "github.com/TimurStash/jwt-go"
	"net/http"
)





func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()
	token, err := findToken(req, authBackend)

	if err == nil && token.Valid && !authBackend.IsInBlacklist(token.Raw) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}

func findToken(req *http.Request, authBackend *JWTAuthenticationBackend) (*jwt.Token, error){
	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	return token, err
}
