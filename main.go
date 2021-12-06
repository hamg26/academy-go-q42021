package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hamg26/academy-go-q42021/config"
	clients "github.com/hamg26/academy-go-q42021/infrastructure/clients"
	"github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	"github.com/hamg26/academy-go-q42021/infrastructure/router"
	"github.com/hamg26/academy-go-q42021/registry"

	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	if config.C.Logging != true {
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	mycsv := datastore.NewCSV(config.C.CSV.Path)
	api := clients.NewPokeApiClient(config.C.API.BaseURL, client)

	r := registry.NewRegistry(mycsv, api)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatal(err)
	}
}
