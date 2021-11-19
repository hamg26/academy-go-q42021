package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/labstack/echo"

	"github.com/hamg26/academy-go-q42021/config"
	"github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	"github.com/hamg26/academy-go-q42021/infrastructure/router"
	"github.com/hamg26/academy-go-q42021/registry"
)

func main() {
	config.ReadConfig()

	err, mycsv := datastore.NewCSV()

	if err != nil {
		panic(err)
	}

	defer mycsv.Close()

	r := registry.NewRegistry(mycsv)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	if config.C.Logging != true {
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)
	}

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server listening at http://localhost" + ":" + config.C.Server.Address)
}
