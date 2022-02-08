package controller

type pokemonController interface {
	GetPokemons(c Context) error
	GetPokemon(c Context) error
	GetPokemonDetails(c Context) error
}

// Interface that defines the methods a controller should implement
type AppController interface {
	pokemonController
}
