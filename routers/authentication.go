package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/TimurStash/gochat/controllers"
	"github.com/TimurStash/gochat/core/authentication"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.Handle("/token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.IsNotAuthenticated),
			negroni.HandlerFunc(controllers.Login),
		)).Methods("POST")

	router.Handle("/signup",
		negroni.New(
			negroni.HandlerFunc(authentication.IsNotAuthenticated),
			negroni.HandlerFunc(controllers.SignUp),
		)).Methods("POST")

	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")
	return router
}