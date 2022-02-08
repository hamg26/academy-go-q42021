package testutil

import (
	"github.com/hamg26/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/mock"
)

func GetPokemons() []*model.Pokemon {
	return []*model.Pokemon{
		{Id: uint64(1), Name: "name1", Type: "type1"},
		{Id: uint64(2), Name: "name2", Type: "type2"},
	}
}

func GetPokemonDetails() *model.PokemonDetails {
	pokemonType := model.PokemonType{Name: "name1", URL: "url1"}
	pokemonTypeSlot := model.PokemonTypeSlot{Slot: 1, Type: pokemonType}
	pokemonTypes := []model.PokemonTypeSlot{pokemonTypeSlot}

	return &model.PokemonDetails{
		Id:    uint64(1),
		Name:  "name1",
		Types: pokemonTypes,
	}
}

type PokemonInteractor struct {
	mock.Mock
}

func (pi *PokemonInteractor) GetOne(id uint64) (error, *model.Pokemon) {
	args := pi.Called(id)
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).(*model.Pokemon)
	}
	return args.Error(1), nil
}

func (pi *PokemonInteractor) GetOneDetails(id string) (error, *model.PokemonDetails) {
	args := pi.Called(id)
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).(*model.PokemonDetails)
	}
	return args.Error(1), nil
}

func (pi *PokemonInteractor) SavePokemon(p *model.PokemonDetails) error {
	args := pi.Called(p)
	return args.Error(0)
}

func (pi *PokemonInteractor) GetAll() (error, []*model.Pokemon) {
	args := pi.Called()
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).([]*model.Pokemon)
	}
	return args.Error(1), nil
}
