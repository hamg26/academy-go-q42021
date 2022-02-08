package registry

import (
	datatstore "github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	controller "github.com/hamg26/academy-go-q42021/interface/controllers"
)

type registry struct {
	mycsv *datatstore.MyCSV
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(mycsv *datatstore.MyCSV) Registry {
	return &registry{mycsv}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
