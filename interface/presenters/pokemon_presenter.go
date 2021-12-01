package presenter

import (
	"sort"
	"strings"

	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/presenter"
)

type pokemonPresenter struct {
}

/*
Returns a new instance of the Pokemon presenter
Specific implementation of endpoint responses using an interactor
*/
func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

/*
Formats the response with all the pokemons (no details)
*/
func (pp *pokemonPresenter) ResponsePokemons(ps []*model.Pokemon) []*model.Pokemon {
	for _, p := range ps {
		pp.ResponsePokemon(p)
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Id < ps[j].Id
	})
	return ps
}

/*
Formats the response of a single pokemon (no details)
*/
func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	if p != nil {
		p.Name = strings.Title(strings.ToLower(p.Name))
		p.Type = strings.Title(strings.ToLower(p.Type))
	}
	return p
}

/*
Formats the response of a single pokemon (with details)
*/
func (pp *pokemonPresenter) ResponsePokemonDetails(p *model.PokemonDetails) *model.PokemonDetails {
	if p != nil {
		p.Name = strings.Title(strings.ToLower(p.Name))
		for i, typeData := range p.Types {
			p.Types[i].Type.Name = strings.Title(strings.ToLower(typeData.Type.Name))
		}
	}
	return p
}
