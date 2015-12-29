package routers

import(
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/TimurStash/gochat/core/authentication"
	"github.com/TimurStash/gochat/socket"
)

func SetWsSocketRoutes(router *mux.Router) *mux.Router {
	router.Handle("/ws/",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(socket.ServeWs),
		))
	return router
}
