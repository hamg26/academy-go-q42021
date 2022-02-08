package presenter

import "github.com/hamg26/academy-go-q42021/domain/model"

// Pokemon presenter interface
// Defines the methods available to use by the interactor
type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
	ResponsePokemonDetails(p *model.PokemonDetails) *model.PokemonDetails
}
