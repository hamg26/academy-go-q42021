package presenter

import "github.com/hamg26/academy-go-q42021/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
}
