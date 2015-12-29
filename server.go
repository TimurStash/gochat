package main
import (
	"github.com/TimurStash/gochat/routers"
	"github.com/TimurStash/gochat/settings"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/rs/cors"
	"github.com/TimurStash/gochat/socket"
)


func main() {
	go socket.RunHub()
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization"},
	}).Handler(n)
	http.ListenAndServe(":5000", handler)
}