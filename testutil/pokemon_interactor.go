package testutil

import (
	"github.com/hamg26/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/mock"
)

// Returns a list of Pokemon for testing purposes
func GetPokemons() []*model.Pokemon {
	return []*model.Pokemon{
		{Id: 1, Name: "name1", Type: "type1"},
		{Id: 2, Name: "name2", Type: "type2"},
	}
}

// Returns an instance PokemonDetails for testing purposes
func GetPokemonDetails() *model.PokemonDetails {
	pokemonType := model.PokemonType{Name: "name1", URL: "url1"}
	pokemonTypeSlot := model.PokemonTypeSlot{Slot: 1, Type: pokemonType}
	pokemonTypes := []model.PokemonTypeSlot{pokemonTypeSlot}

	return &model.PokemonDetails{
		Id:    1,
		Name:  "name1",
		Types: pokemonTypes,
	}
}

// Mocked PokemonInteractor
type PokemonInteractor struct {
	mock.Mock
}

// Mocked PokemonInteractor.GetOne
func (pi *PokemonInteractor) GetOne(id int) (error, *model.Pokemon) {
	args := pi.Called(id)
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).(*model.Pokemon)
	}
	return args.Error(1), nil
}

// Mocked PokemonInteractor.GetOneDetails
func (pi *PokemonInteractor) GetOneDetails(id int) (error, *model.PokemonDetails) {
	args := pi.Called(id)
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).(*model.PokemonDetails)
	}
	return args.Error(1), nil
}

// Mocked PokemonInteractor.SavePokemon
func (pi *PokemonInteractor) SavePokemon(p *model.PokemonDetails) error {
	args := pi.Called(p)
	return args.Error(0)
}

// Mocked PokemonInteractor.GetAll
func (pi *PokemonInteractor) GetAll() (error, []*model.Pokemon) {
	args := pi.Called()
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).([]*model.Pokemon)
	}
	return args.Error(1), nil
}

// Mocked PokemonInteractor.GetAllConcurrent
func (pi *PokemonInteractor) GetAllConcurrent(string, int, int) (error, []*model.Pokemon) {
	args := pi.Called()
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).([]*model.Pokemon)
	}
	return args.Error(1), nil
}
