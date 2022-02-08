package registry

import (
	controller "github.com/hamg26/academy-go-q42021/interface/controllers"
	ip "github.com/hamg26/academy-go-q42021/interface/presenters"
	ir "github.com/hamg26/academy-go-q42021/interface/repository"
	"github.com/hamg26/academy-go-q42021/usecase/interactor"
	pp "github.com/hamg26/academy-go-q42021/usecase/presenter"
	pr "github.com/hamg26/academy-go-q42021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() pr.PokemonRepository {
	return ir.NewPokemonRepository(r.mycsv)
}

func (r *registry) NewPokemonPresenter() pp.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
