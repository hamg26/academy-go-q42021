package repository

import "github.com/hamg26/academy-go-q42021/domain/model"

// Pokemon repository interface
// Defines the methods available to use by the interactor
type PokemonRepository interface {
	FindAll() (error, []*model.Pokemon)
	FindOne(id uint64) (error, *model.Pokemon)
	FindOneDetails(id string) (error, *model.PokemonDetails)
	SavePokemon(*model.PokemonDetails) error
}
