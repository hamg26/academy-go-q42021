package presenter

import (
	"strings"

	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/presenter"
)

type pokemonPresenter struct {
}

func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(ps []*model.Pokemon) []*model.Pokemon {
	for _, p := range ps {
		p.Name = strings.Title(strings.ToLower(p.Name))
		p.Type = strings.Title(strings.ToLower(p.Type))
	}
	return ps
}

func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	if p != nil {
		p.Name = strings.Title(strings.ToLower(p.Name))
		p.Type = strings.Title(strings.ToLower(p.Type))
	}
	return p
}
